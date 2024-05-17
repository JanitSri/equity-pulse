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

	er := model.NewEquityRequestBuilder().Ticker("TSLA").Build()

	c := &http.Client{}
	y := net.NewYahooFinanceDataProvider(c)

	e := service.NewEquityService(y)
	r, err := e.EndOfDayStockPrices(er)

	if err != nil {
		fmt.Println("<<<ERROR>>>", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", r)
}
