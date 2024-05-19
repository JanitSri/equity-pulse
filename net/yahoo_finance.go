package net

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/model/yahoo"
	"github.com/JanitSri/equity-pulse/util"
	"golang.org/x/net/publicsuffix"
)

const (
	yFDefaultScheme = "https"
	yFBaseURL       = "finance.yahoo.com"
	yFQuery1URL     = "query1." + yFBaseURL
	yFQuery2URL     = "query2." + yFBaseURL
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
		"lang":      {"en-CA"},
		"region":    {"CA"},
		"modules":   {"summaryProfile,details"},
	}

	u, _ := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), fmt.Sprintf("/v10/finance/quoteSummary/%s", ticker), p)

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
	h.Set(util.HostHeader, yFQuery2URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

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

func (y *YahooFinanceDataProvider) RetrieveStockPrices(ticker string, start, end time.Time, interval model.Interval) (*model.EndOfDayPrices, error) {
	ic := model.NewIntervalConverter(interval)
	i := ic.ConvertToYahooFinanceInterval()

	if i == model.InvalidInterval {
		return nil, fmt.Errorf("invalid interval: %s", interval)
	}

	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-CA"},
		"region":    {"CA"},
		"interval":  {string(i)},
		"period1":   {strconv.FormatInt(start.Unix(), 10)},
		"period2":   {strconv.FormatInt(end.Unix(), 10)},
	}

	u, _ := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), fmt.Sprintf("/v8/finance/chart/%s", ticker), p)

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery2URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	s := &yahoo.EndOfDay{}
	util.FetchAndDecode(y.client, u, http.MethodGet, h, s)

	if len(s.Chart.Result) == 0 {
		return nil, fmt.Errorf("empty result for EndOfDay")
	}

	t := s.Chart.Result[0].Timestamp
	q := s.Chart.Result[0].Indicators.Quote[0]
	ac := s.Chart.Result[0].Indicators.Adjclose[0]

	eod := make(model.EndOfDayPrices, len(t))

	for i := range t {
		s := model.NewStockPriceBuilder().Time(t[i].Time).High(q.High[i]).Low(q.Low[i]).Open(q.Open[i]).Close(ac.Close[i]).Volume(q.Volume[i]).Build()
		eod[i] = s
	}

	return &eod, nil
}

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error) {
	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-CA"},
		"region":    {"CA"},
	}

	u, _ := util.BuildURL(fmt.Sprintf("https://%s", yFQuery1URL), fmt.Sprintf("/v1/finance/quoteType/%s", ticker), p)

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery1URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

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
	u, _ := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), "/v1/test/getcrumb", nil)

	h := http.Header{}
	h.Set(util.AcceptHeader, "*/*")
	h.Set(util.HostHeader, yFQuery1URL)

	var c util.TextBodyResponse
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, &c); err != nil {
		return "", err
	}

	return c.Body, nil
}

func (y *YahooFinanceDataProvider) verifyYahooFinanceCookie(targetUrl *url.URL) error {
	cu, _ := util.BuildURL(fmt.Sprintf("https://%s", yFBaseURL), "", nil)
	if err := util.VerifyCookie(y.client, targetUrl, cu); err != nil {
		return err
	}

	return nil
}
