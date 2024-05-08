package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/service"
)

func main() {
	// cmd.Execute()

	er := model.NewEquityRequestBuilder().Ticker("AAPL").Build()

	cj, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar: cj,
	}
	y := net.NewYahooFinanceDataProvider(c)

	e := service.NewEquityService(y)
	t, _ := e.StockTickerInfo(er)

	fmt.Printf("%+v\n", t)
}
