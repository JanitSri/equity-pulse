package service

import (
	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
)

type EquityService struct {
	stockDataProvider net.StockDataProvider
}

func NewEquityService(stockDataProvider net.StockDataProvider) *EquityService {
	return &EquityService{
		stockDataProvider,
	}
}

func (e *EquityService) StockNews(er *model.EquityRequest) *model.News {
	return e.stockDataProvider.RetrieveStockNews(er.Ticker())
}
func (e *EquityService) CompanyProfile(er *model.EquityRequest) *model.CompanyProfile {
	return e.stockDataProvider.RetrieveCompanyProfile(er.Ticker())
}
func (e *EquityService) StockStatistics(er *model.EquityRequest) *model.StockStatistics {
	return e.stockDataProvider.RetrieveStockStatistics(er.Ticker())

}
func (e *EquityService) EndOfDayStockPrices(er *model.EquityRequest) *model.EndOfDayPrices {
	return e.stockDataProvider.RetrieveEndOfDayStockPrices(er.Ticker())

}
func (e *EquityService) StockTickerInfo(er *model.EquityRequest) *model.TickerInfo {
	return e.stockDataProvider.RetrieveStockTickerInfo(er.Ticker())
}
