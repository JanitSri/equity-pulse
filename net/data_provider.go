package net

import "github.com/JanitSri/equity-pulse/model"

type StockDataProvider interface {
	RetrieveStockNews(ticker string) *model.News
	RetrieveCompanyProfile(ticker string) *model.CompanyProfile
	RetrieveStockStatistics(ticker string) *model.StockStatistics
	RetrieveEndOfDayStockPrices(ticker string) *model.EndOfDayPrices
	RetrieveStockTickerInfo(ticker string) *model.TickerInfo
}
