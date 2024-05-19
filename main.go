package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"github.com/JanitSri/equity-pulse/service"
)

// TODO: add remove fmt and add logging

func main() {
	// cmd.Execute()

	er := model.NewEquityRequestBuilder().Ticker("AAPL").Build()

	c := &http.Client{}
	y := net.NewYahooFinanceDataProvider(c)

	e := service.NewEquityService(y)

	// start, _ := time.Parse(time.RFC3339, "2024-05-13T00:00:00-04:00")
	// end, _ := time.Parse(time.RFC3339, "2024-05-18T00:00:00-04:00")
	// r, err := e.EndOfDayStockPrices(er, start, end)

	// r, err := e.CompanyProfile(er)

	// r, err := e.StockTickerInfo(er)

	r, err := e.StockStatistics(er)

	if err != nil {
		fmt.Println("<<<ERROR>>>", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", r)
}
