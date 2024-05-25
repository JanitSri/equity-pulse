package service

import (
	"time"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"go.uber.org/zap"
)

type EquityService struct {
	stockDataProvider net.StockDataProvider
}

func NewEquityService(s net.StockDataProvider) *EquityService {
	return &EquityService{
		stockDataProvider: s,
	}
}

func (e *EquityService) StockNews(er *model.EquityRequest) (*model.News, error) {
	zap.L().Sugar().Infof("retrieving stock news for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveStockNews(er.Ticker())
}
func (e *EquityService) CompanyProfile(er *model.EquityRequest) (*model.CompanyProfile, error) {
	zap.L().Sugar().Infof("retrieving company profile for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveCompanyProfile(er.Ticker())
}
func (e *EquityService) StockStatistics(er *model.EquityRequest) (*model.StockStatistics, error) {
	zap.L().Sugar().Infof("retrieving stock statistics for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveStockStatistics(er.Ticker())

}
func (e *EquityService) EndOfDayStockPrices(er *model.EquityRequest, start, end time.Time) (*model.EndOfDayPrices, error) {
	zap.L().Sugar().Infof("retrieving end of day stock prices for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveStockPrices(er.Ticker(), start, end, model.Interval1d)

}
func (e *EquityService) StockTickerInfo(er *model.EquityRequest) (*model.TickerInfo, error) {
	zap.L().Sugar().Infof("retrieving stock ticker info for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveStockTickerInfo(er.Ticker())
}
