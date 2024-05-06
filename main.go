package main

import (
	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/service"
)

func main() {
	// cmd.Execute()

	er := model.NewEquityRequestBuilder().Ticker("AAPL").Build()

	y := net.NewYahooFinanceDataProvider()
	e := service.NewEquityService(y)

	e.StockTickerInfo(er)
}
