package net

import (
	"github.com/JanitSri/equity-pulse/model"
)

const (
	userAgentHeader = "User-Agent"
	acceptHeader    = "Accept"
	hostHeader      = "Host"
)

type StockDataProvider interface {
	RetrieveStockNews(ticker string) (*model.News, error)
	RetrieveCompanyProfile(ticker string) (*model.CompanyProfile, error)
	RetrieveStockStatistics(ticker string) (*model.StockStatistics, error)
	RetrieveEndOfDayStockPrices(ticker string) (*model.EndOfDayPrices, error)
	RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error)
}
