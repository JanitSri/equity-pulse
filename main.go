package main

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
	"sync"
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

	var wg sync.WaitGroup

	start := time.Now()

	results := make(chan Result, 3)

	handleFunc := func(name string, fn func(*model.EquityRequest) (interface{}, error)) {
		defer wg.Done()
		result, err := fn(er)
		results <- Result{Name: name, Result: result, Err: err}
	}

	wg.Add(1)
	go handleFunc("EndOfDayStockPrices", func(er *model.EquityRequest) (interface{}, error) {
		start, _ := time.Parse(time.RFC3339, "2024-05-13T00:00:00-04:00")
		end, _ := time.Parse(time.RFC3339, "2024-05-18T00:00:00-04:00")
		return e.EndOfDayStockPrices(er, start, end)
	})

	wg.Add(1)
	go handleFunc("CompanyProfile", func(er *model.EquityRequest) (interface{}, error) {
		return e.CompanyProfile(er)
	})

	wg.Add(1)
	go handleFunc("StockTickerInfo", func(er *model.EquityRequest) (interface{}, error) {
		return e.StockTickerInfo(er)
	})

	wg.Add(1)
	go handleFunc("StockNews", func(er *model.EquityRequest) (interface{}, error) {
		return e.StockNews(er)
	})

	go func(s time.Time) {
		wg.Wait()
		totalElapsed := time.Since(s)
		zap.L().Sugar().Infof("Total time for all goroutines: %s", totalElapsed)
		close(results)
	}(start)

	for res := range results {
		if res.Err != nil {
			zap.L().Sugar().Errorf("%s: something went wrong: %v", res.Name, res.Err)
		} else {
			zap.L().Sugar().Debugf("%s: %+v", res.Name, res.Result)
		}
	}
}

// RUN TIMES (BEFORE PIPELINING):
// 		1.  1.852192458s
// 		2.  1.388701917s
//      3.  985.604791ms
// 		4.  1.28646025s
// 		5.  1.405365292s
// 		6.  1.423375542s
// 		7.  1.617825417s
// 		8.  1.635723666s
// 		9.  1.318731083s
// 		10. 1.449131792s
//
//  AVG TIME: 1.436 seconds
//
//
// TODO: add redis pipelining commands
//
