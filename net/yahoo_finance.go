package net

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/model/yahoo"
	"github.com/JanitSri/equity-pulse/util"
	"golang.org/x/net/publicsuffix"
)

const (
	yFQuery1URL = "https://query1.finance.yahoo.com"
	yFQuery2URL = "https://query2.finance.yahoo.com"
)

type YahooFinanceDataProvider struct {
	client *http.Client
	cache  cache
}

func NewYahooFinanceDataProvider(c *http.Client) *YahooFinanceDataProvider {
	// TODO: implement cache
	o := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	cj, _ := cookiejar.New(&o)
	c.Jar = cj

	return &YahooFinanceDataProvider{
		client: c,
	}
}

// TODO: handle errors -> return response and error and let upstream handle it

func (y *YahooFinanceDataProvider) RetrieveStockNews(ticker string) (*model.News, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveCompanyProfile(ticker string) (*model.CompanyProfile, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveStockStatistics(ticker string) (*model.StockStatistics, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveEndOfDayStockPrices(ticker string) (*model.EndOfDayPrices, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error) {
	p := map[string][]string{
		"formatted":  {"true"},
		"amp;lang":   {"en-US"},
		"amp;region": {"US"},
	}

	u, _ := util.BuildURL(yFQuery1URL, fmt.Sprintf("/v1/finance/quoteType/%s", ticker), p)

	h := http.Header{}
	h.Set(util.HostHeader, "query1.finance.yahoo.com")
	h.Set(util.AcceptHeader, "application/json")

	q := &yahoo.QuoteInfo{}
	err := util.FetchAndDecode(y.client, u, http.MethodGet, h, q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if len(q.QuoteType.Result) == 0 {
		return nil, fmt.Errorf("empty result for QuoteType")
	}

	r := q.QuoteType.Result[0]
	t := model.NewTickerInfoBuilder().Symbol(r.Symbol).Exchange(r.Exchange).QuoteType(r.QuoteType).LongName(r.LongName).ShortName(r.ShortName).Build()

	return t, nil
}

func (y *YahooFinanceDataProvider) getCrumb() (string, error) {
	u, _ := util.BuildURL(yFQuery2URL, "/v1/test/getcrumb", nil)

	h := http.Header{}
	h.Set(util.AcceptHeader, "*/*")
	h.Set(util.HostHeader, "query1.finance.yahoo.com")

	var c util.TextBodyResponse
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, &c); err != nil {
		return "", err
	}

	return c.Body, nil
}

func (y *YahooFinanceDataProvider) verifyYahooFinanceCookie(targetUrl *url.URL) error {
	cu, _ := util.BuildURL("https://finance.yahoo.com", "", nil)
	if err := util.VerifyCookie(y.client, targetUrl, cu); err != nil {
		return err
	}

	return nil
}
