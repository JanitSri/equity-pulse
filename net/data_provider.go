package net

import (
	"time"

	"github.com/JanitSri/equity-pulse/model"
)

type StockDataProvider interface {
	RetrieveStockNews(ticker string) (model.ArticleIds, error)
	RetrieveArticle(id string) (*model.ArticleBuilder, error)
	RetrieveCompanyProfile(ticker string) (*model.CompanyProfileBuilder, error)
	RetrieveStockStatistics(ticker string) (*model.StockStatisticsBuilder, error)
	RetrieveStockPrices(ticker string, start, end time.Time, interval model.Interval) (*model.EndOfDayPrices, error)
	RetrieveStockTickerInfo(ticker string) (*model.TickerInfoBuilder, error)
}
