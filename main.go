package main

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"time"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/service"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/publicsuffix"
)

func createLogger() *zap.Logger {
	env := os.Getenv("APP_ENV")

	c := zap.NewProductionEncoderConfig()
	c.EncodeLevel = zapcore.CapitalLevelEncoder
	c.TimeKey = "timestamp"
	c.EncodeTime = zapcore.ISO8601TimeEncoder
	c.MessageKey = "message"

	cE := zapcore.NewJSONEncoder(c)
	s := zapcore.AddSync(os.Stdout)

	var core zapcore.Core
	if strings.EqualFold(env, "PROD") {
		core = zapcore.NewCore(cE, s, zap.NewAtomicLevelAt(zap.InfoLevel))
	} else {
		core = zapcore.NewCore(cE, s, zap.NewAtomicLevelAt(zap.DebugLevel))
	}

	return zap.New(core)
}

func init() {
	zap.ReplaceGlobals(createLogger())
}

type Result struct {
	Name   string
	Result interface{}
	Err    error
}

func main() {
	// cmd.Execute()

	er := model.NewEquityRequestBuilder().Ticker("AAPL").Build()

	ctx := context.Background()

	opts := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	rC := net.NewRedisCache(opts)

	pong, err := rC.CheckConnection(ctx)
	if err != nil {
		zap.L().Sugar().Errorf("could not connect to redis: %s", err)
	} else {
		zap.L().Sugar().Infof("connected to redis: %s", pong)
	}

	cs := service.NewCacheService(rC)

	c := &http.Client{}
	o := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	cj, _ := cookiejar.New(&o)
	c.Jar = cj

	y := net.NewYahooFinanceDataProvider(c)

	e := service.NewEquityService(y, cs)

	startTime := time.Now()

	sn, err := e.StockNews(er)
	if err != nil {
		zap.L().Sugar().Errorf("StockNews err: %s", err)
	}
	zap.L().Sugar().Infof("StockNews res: %s", *sn)

	start, _ := time.Parse(time.RFC3339, "2024-05-13T00:00:00-04:00")
	end, _ := time.Parse(time.RFC3339, "2024-05-18T00:00:00-04:00")
	eod, err := e.EndOfDayStockPrices(er, start, end)
	if err != nil {
		zap.L().Sugar().Errorf("EndOfDayStockPrices err: %s", end)
	}
	zap.L().Sugar().Infof("EndOfDayStockPrices res: %s", *eod)

	cp, err := e.CompanyProfile(er)
	if err != nil {
		zap.L().Sugar().Errorf("CompanyProfile err: %s", end)
	}
	zap.L().Sugar().Infof("CompanyProfile res: %s", *cp)

	st, err := e.StockTickerInfo(er)
	if err != nil {
		zap.L().Sugar().Errorf("StockTickerInfo err: %s", end)
	}
	zap.L().Sugar().Infof("StockTickerInfo res: %s", *st)

	zap.L().Sugar().Infof("Total time for all goroutines: %s", time.Since(startTime))
}
