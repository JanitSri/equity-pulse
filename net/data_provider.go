package net

import (
	"time"

	"github.com/JanitSri/equity-pulse/model"
)

type StockDataProvider interface {
	RetrieveStockNews(ticker string) (*model.News, error)
	RetrieveCompanyProfile(ticker string) (*model.CompanyProfile, error)
	RetrieveStockStatistics(ticker string) (*model.StockStatistics, error)
	RetrieveStockPrices(ticker string, start, end time.Time, interval model.Interval) (*model.EndOfDayPrices, error)
	RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error)
}
