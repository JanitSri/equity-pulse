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
	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-US"},
		"region":    {"US"},
		"modules":   {"summaryProfile,details"},
	}

	u, _ := util.BuildURL(yFQuery2URL, fmt.Sprintf("/v10/finance/quoteSummary/%s", ticker), p)

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	c, err := y.getCrumb()
	if err != nil {
		return nil, err
	}

	cr := url.Values{}
	cr.Add("crumb", c)
	u.RawQuery += fmt.Sprintf("&%s", cr.Encode())

	h := http.Header{}
	h.Set(util.HostHeader, "qquery2.finance.yahoo.com")
	h.Set(util.AcceptHeader, "application/json")

	yc := &yahoo.CompanyProfile{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, yc); err != nil {
		return nil, err
	}

	if len(yc.QuoteSummary.Result) == 0 {
		return nil, fmt.Errorf("empty result for Company Profile")
	}

	r := yc.QuoteSummary.Result[0]
	cp := model.NewCompanyProfileBuilder().Address(r.SummaryProfile.Address1).City(r.SummaryProfile.City).State(r.SummaryProfile.State).Zip(r.SummaryProfile.Zip).Country(r.SummaryProfile.Country).Phone(r.SummaryProfile.Phone).Website(r.SummaryProfile.Website).Industry(r.SummaryProfile.Industry).Sector(r.SummaryProfile.Sector).LongBusinessSummary(r.SummaryProfile.LongBusinessSummary).FullTimeEmployee(uint32(r.SummaryProfile.FullTimeEmployees)).InvestorRelationWebsite(r.SummaryProfile.IrWebsite).Build()

	return cp, nil
}

func (y *YahooFinanceDataProvider) RetrieveStockStatistics(ticker string) (*model.StockStatistics, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveEndOfDayStockPrices(ticker string) (*model.EndOfDayPrices, error) {
	panic("not implemented") // TODO: Implement
}

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error) {
	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-US"},
		"region":    {"US"},
	}

	u, _ := util.BuildURL(yFQuery1URL, fmt.Sprintf("/v1/finance/quoteType/%s", ticker), p)

	h := http.Header{}
	h.Set(util.HostHeader, "query1.finance.yahoo.com")
	h.Set(util.AcceptHeader, "application/json")

	q := &yahoo.QuoteInfo{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, q); err != nil {
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
