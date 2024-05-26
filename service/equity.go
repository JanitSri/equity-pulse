package service

import (
	"fmt"
	"time"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	"go.uber.org/zap"
)

const (
	stockNewsArticleKey    = "%s-stockNewsArticle"
	companyProfileKey      = "%s-CompanyProfile"
	stockStatisticsKey     = "%s-StockStatistics"
	endOfDayStockPricesKey = "%s-EndOfDayStockPrices"
	stockTickerInfoKey     = "%s-StockTickerInfo"

	dailyTTL     = time.Hour * 24
	monthlyTTL   = time.Hour * 24 * 30
	quarterlyTTL = time.Hour * 24 * 90
)

type EquityService struct {
	stockDataProvider net.StockDataProvider
	cache             *CacheService
}

func NewEquityService(s net.StockDataProvider, c *CacheService) *EquityService {
	return &EquityService{
		stockDataProvider: s,
		cache:             c,
	}
}

func (e *EquityService) StockNews(er *model.EquityRequest) (*model.News, error) {
	zap.L().Sugar().Infof("retrieving stock news for %s %s", er.Name(), er.Ticker())

	ids, err := e.stockDataProvider.RetrieveStockNews(er.Ticker())
	if err != nil {
		return nil, err
	}

	n := make(model.News, 0)
	for _, id := range ids {
		if len(id) > 0 {
			res := &model.ArticleBuilder{}
			if err := e.cache.Get(fmt.Sprintf(stockNewsArticleKey, id), res); err == nil {
				n = append(n, res.Build())
				continue
			}

			ab, err := e.stockDataProvider.RetrieveArticle(id)
			if err != nil {
				zap.L().Sugar().Errorf("cannot retrieve article id %s: %s", id, err)
				continue
			}
			e.cache.Set(fmt.Sprintf(stockNewsArticleKey, id), ab, monthlyTTL)

			n = append(n, ab.Build())
		}
	}

	return &n, nil
}

func (e *EquityService) CompanyProfile(er *model.EquityRequest) (*model.CompanyProfile, error) {
	zap.L().Sugar().Infof("retrieving company profile for %s %s", er.Name(), er.Ticker())

	res := &model.CompanyProfileBuilder{}
	if err := e.cache.Get(fmt.Sprintf(companyProfileKey, er.Ticker()), res); err == nil {
		return res.Build(), nil
	}

	cp, err := e.stockDataProvider.RetrieveCompanyProfile(er.Ticker())
	if err != nil {
		return nil, err
	}
	e.cache.Set(fmt.Sprintf(companyProfileKey, er.Ticker()), cp, monthlyTTL)

	return cp.Build(), nil
}

func (e *EquityService) StockStatistics(er *model.EquityRequest) (*model.StockStatistics, error) {
	zap.L().Sugar().Infof("retrieving stock statistics for %s %s", er.Name(), er.Ticker())

	res := &model.StockStatisticsBuilder{}
	if err := e.cache.Get(fmt.Sprintf(stockStatisticsKey, er.Ticker()), res); err == nil {
		return res.Build(), nil
	}

	s, err := e.stockDataProvider.RetrieveStockStatistics(er.Ticker())
	if err != nil {
		return nil, err
	}
	e.cache.Set(fmt.Sprintf(stockStatisticsKey, er.Ticker()), s, monthlyTTL)

	return s.Build(), nil
}

func (e *EquityService) EndOfDayStockPrices(er *model.EquityRequest, start, end time.Time) (*model.EndOfDayPrices, error) {
	zap.L().Sugar().Infof("retrieving end of day stock prices for %s %s", er.Name(), er.Ticker())
	return e.stockDataProvider.RetrieveStockPrices(er.Ticker(), start, end, model.Interval1d)

}

func (e *EquityService) StockTickerInfo(er *model.EquityRequest) (*model.TickerInfo, error) {
	zap.L().Sugar().Infof("retrieving stock ticker info for %s %s", er.Name(), er.Ticker())

	res := &model.TickerInfoBuilder{}
	if err := e.cache.Get(fmt.Sprintf(stockTickerInfoKey, er.Ticker()), res); err == nil {
		return res.Build(), nil
	}

	t, err := e.stockDataProvider.RetrieveStockTickerInfo(er.Ticker())
	if err != nil {
		return nil, err
	}
	e.cache.Set(fmt.Sprintf(stockTickerInfoKey, er.Ticker()), t, monthlyTTL)

	return t.Build(), nil
}
