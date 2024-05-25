package main

import (
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/service"
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

func main() {
	// cmd.Execute()

	er := model.NewEquityRequestBuilder().Ticker("AAPL").Build()

	c := &http.Client{}
	o := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	cj, _ := cookiejar.New(&o)
	c.Jar = cj

	y := net.NewYahooFinanceDataProvider(c)

	e := service.NewEquityService(y)

	// start, _ := time.Parse(time.RFC3339, "2024-05-13T00:00:00-04:00")
	// end, _ := time.Parse(time.RFC3339, "2024-05-18T00:00:00-04:00")
	// r, err := e.EndOfDayStockPrices(er, start, end)

	// r, err := e.CompanyProfile(er)

	r, err := e.StockTickerInfo(er)

	// r, err := e.StockStatistics(er)

	// r, err := e.StockNews(er)

	if err != nil {
		zap.L().Sugar().Fatal("something went wrong", zap.Error(err))
	}

	zap.L().Sugar().Debugf("%+v", *r)
}
