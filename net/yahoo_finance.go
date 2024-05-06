package net

import "github.com/JanitSri/equity-pulse/model"

type YahooFinanceDataProvider struct{}

func (y *YahooFinanceDataProvider) RetrieveStockNews(ticker string) model.News {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveCompanyProfile(ticker string) model.CompanyProfile {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveStockStatistics(ticker string) model.StockStatistics {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveEndOfDayStockPrices(ticker string) model.EndOfDayPrices {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) model.TickerInfo {
	panic("not implemented") // TODO: Implement
}
