package net

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/model/yahoo"
	"github.com/JanitSri/equity-pulse/util"
	"go.uber.org/zap"
)

const (
	yFDefaultScheme = "https"
	yFBaseURL       = "finance.yahoo.com"
	yFQuery1URL     = "query1." + yFBaseURL
	yFQuery2URL     = "query2." + yFBaseURL
	yFMediaURL      = "ncp-gw-finance.media.yahoo.com"
)

type YahooFinanceDataProvider struct {
	client *http.Client
}

func NewYahooFinanceDataProvider(c *http.Client) *YahooFinanceDataProvider {
	return &YahooFinanceDataProvider{
		client: c,
	}
}

func (y *YahooFinanceDataProvider) RetrieveStockNews(ticker string) (model.ArticleIds, error) {
	zap.L().Sugar().Debugf("yahoo finance retrieving stock news for %s", ticker)

	p := map[string][]string{
		"count":        {"250"},
		"namespace":    {"finance"},
		"snippetCount": {"20"},
		"id":           {"tickers-all-stream"},
		"version":      {"v2"},
		"device":       {"desktop"},
		"site":         {"finance"},
		"s":            {ticker},
		"lang":         {"en-CA"},
		"region":       {"CA"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFMediaURL), "/api/v2/gql/stream_view", p)
	if err != nil {
		return nil, err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFMediaURL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	cn := &yahoo.CompanyNews{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, cn); err != nil {
		return nil, err
	}

	if len(cn.Data.TickerStream.Pagination.Uuids) == 0 {
		return nil, fmt.Errorf("empty result for company news")
	}

	r := cn.Data.TickerStream.Pagination.Uuids

	re, err := regexp.Compile(`"id":"([a-f0-9-]+)"`)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllStringSubmatch(r, -1)

	a := model.ArticleIds{}
	m := make(map[string]bool)
	for _, match := range matches {
		if _, ok := m[match[1]]; ok {
			continue
		}
		a = append(a, match[1])
		m[match[1]] = true
	}

	return a, nil
}

func (y *YahooFinanceDataProvider) RetrieveArticle(id string) (*model.ArticleBuilder, error) {
	zap.L().Sugar().Debugf("Retrieving yahoo finance article id %s", id)

	p := map[string][]string{
		"device": {"desktop"},
		"site":   {"finance"},
		"lang":   {"en-CA"},
		"region": {"CA"},
		"bot":    {"0"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFBaseURL), "/caas/content/article", p)
	if err != nil {
		return nil, err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFMediaURL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	cr := u.Query()
	cr.Add("uuid", id)
	u.RawQuery = cr.Encode()

	ya := &yahoo.Article{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, ya); err != nil {
		return nil, err
	}

	if len(ya.Items) == 0 {
		return nil, fmt.Errorf("no data for article uuid %s", id)
	}

	r := ya.Items[0]
	a := &model.ArticleBuilder{
		Title:         r.Schema.Default.Headline,
		Summary:       r.Schema.Default.Description,
		Url:           r.Schema.Default.MainEntityOfPage,
		DatePublished: r.Schema.Default.DatePublished,
		DateUpdated:   r.Schema.Default.DateModified,
		Keywords:      r.Schema.Default.Keywords,
	}

	return a, nil
}

func (y *YahooFinanceDataProvider) RetrieveCompanyProfile(ticker string) (*model.CompanyProfileBuilder, error) {
	zap.L().Sugar().Debugf("yahoo finance retrieving company profile for %s", ticker)

	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-CA"},
		"region":    {"CA"},
		"modules":   {"summaryProfile,details"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), fmt.Sprintf("/v10/finance/quoteSummary/%s", ticker), p)
	if err != nil {
		return nil, err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	c, err := y.getCrumb()
	if err != nil {
		return nil, err
	}

	cr := u.Query()
	cr.Add("crumb", c)
	u.RawQuery = cr.Encode()

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery2URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	yc := &yahoo.CompanyProfile{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, yc); err != nil {
		return nil, err
	}

	if len(yc.QuoteSummary.Result) == 0 {
		return nil, fmt.Errorf("empty result for company profile")
	}

	r := yc.QuoteSummary.Result[0]
	cp := &model.CompanyProfileBuilder{
		Address:                 r.SummaryProfile.Address1,
		City:                    r.SummaryProfile.City,
		State:                   r.SummaryProfile.State,
		Zip:                     r.SummaryProfile.Zip,
		Country:                 r.SummaryProfile.Country,
		Phone:                   r.SummaryProfile.Phone,
		Website:                 r.SummaryProfile.Website,
		Industry:                r.SummaryProfile.Industry,
		Sector:                  r.SummaryProfile.Sector,
		LongBusinessSummary:     r.SummaryProfile.LongBusinessSummary,
		FullTimeEmployee:        uint32(r.SummaryProfile.FullTimeEmployees),
		InvestorRelationWebsite: r.SummaryProfile.IrWebsite,
	}

	return cp, nil
}

func (y *YahooFinanceDataProvider) RetrieveStockStatistics(ticker string) (*model.StockStatisticsBuilder, error) {
	zap.L().Sugar().Debugf("yahoo finance retrieving stock statistics for %s", ticker)

	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-CA"},
		"region":    {"CA"},
		"modules":   {"defaultKeyStatistics,financialData"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), fmt.Sprintf("/v10/finance/quoteSummary/%s", ticker), p)
	if err != nil {
		return nil, err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	c, err := y.getCrumb()
	if err != nil {
		return nil, err
	}

	cr := u.Query()
	cr.Add("crumb", c)
	u.RawQuery = cr.Encode()

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery2URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	cs := &yahoo.CompanyStatistics{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, cs); err != nil {
		return nil, err
	}

	if len(cs.QuoteSummary.Result) == 0 {
		return nil, fmt.Errorf("empty result for company statistics")
	}

	r := cs.QuoteSummary.Result[0].DefaultKeyStatistics
	s := &model.StockStatisticsBuilder{
		FiftyTwoWeekChange:       r.FiftyTwoWeekChange.Fmt,
		Beta:                     r.Beta.Fmt,
		BookValue:                r.BookValue.Fmt,
		EnterpriseToEbitda:       r.EnterpriseToEbitda.Fmt,
		EnterpriseToRevenue:      r.EnterpriseToRevenue.Fmt,
		EnterpriseValue:          r.EnterpriseValue.Fmt,
		FiveYearAverageReturn:    r.FiveYearAverageReturn.Fmt,
		FloatShares:              r.FloatShares.Fmt,
		ForwardEps:               r.ForwardEps.Fmt,
		ForwardPE:                r.ForwardPE.Fmt,
		HeldPercentInsiders:      r.HeldPercentInsiders.Fmt,
		HeldPercentInstitutions:  r.HeldPercentInstitutions.Fmt,
		ImpliedSharesOutstanding: r.ImpliedSharesOutstanding.Fmt,
		LastDividendDate:         r.LastDividendDate.Fmt,
		LastDividendValue:        r.LastDividendValue.Fmt,
		NetIncomeToCommon:        r.NetIncomeToCommon.Fmt,
		PegRatio:                 r.PegRatio.Fmt,
		PriceToBook:              r.PriceToBook.Fmt,
		ProfitMargins:            r.ProfitMargins.Fmt,
		RevenueQuarterlyGrowth:   r.RevenueQuarterlyGrowth.Fmt,
		SharesOutstanding:        r.SharesOutstanding.Fmt,
		SharesShort:              r.SharesShort.Fmt,
		ShortRatio:               r.ShortRatio.Fmt,
		TotalAssets:              r.TotalAssets.Fmt,
		TrailingEps:              r.TrailingEps.Fmt,
		Yield:                    r.Yield.Fmt,
		YtdReturn:                r.YtdReturn.Fmt,
	}
	return s, nil
}

func (y *YahooFinanceDataProvider) RetrieveStockPrices(ticker string, start, end time.Time, interval model.Interval) (*model.EndOfDayPrices, error) {
	zap.L().Sugar().Debugf("yahoo finance retrieving stock prices for %s from %s to %s in %s interval", ticker, start, end, interval)

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

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), fmt.Sprintf("/v8/finance/chart/%s", ticker), p)
	if err != nil {
		return nil, err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery2URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	s := &yahoo.EndOfDay{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, s); err != nil {
		return nil, err
	}

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

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) (*model.TickerInfoBuilder, error) {
	zap.L().Sugar().Debugf("yahoo finance retrieving stock ticker info for %s", ticker)

	p := map[string][]string{
		"formatted": {"true"},
		"lang":      {"en-CA"},
		"region":    {"CA"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFQuery1URL), fmt.Sprintf("/v1/finance/quoteType/%s", ticker), p)
	if err != nil {
		return nil, err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFQuery1URL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	q := &yahoo.QuoteInfo{}
	if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, q); err != nil {
		return nil, err
	}

	if len(q.QuoteType.Result) == 0 {
		return nil, fmt.Errorf("empty result for ticker info")
	}

	r := q.QuoteType.Result[0]
	t := &model.TickerInfoBuilder{
		Symbol:    r.Symbol,
		Exchange:  r.Exchange,
		QuoteType: r.QuoteType,
		LongName:  r.LongName,
		ShortName: r.ShortName,
	}

	return t, nil
}

func (y *YahooFinanceDataProvider) getCrumb() (string, error) {
	zap.L().Sugar().Debug("yahoo finance retrieving crumb")

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFQuery2URL), "/v1/test/getcrumb", nil)
	if err != nil {
		return "", err
	}

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
	zap.L().Sugar().Debug("yahoo finance retrieving cookie")

	cu, err := util.BuildURL(fmt.Sprintf("https://%s", yFBaseURL), "", nil)
	if err != nil {
		return err
	}
	h := http.Header{}
	h.Set(util.AcceptHeader, "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	if err := util.VerifyCookie(y.client, targetUrl, cu, h); err != nil {
		return err
	}

	return nil
}
