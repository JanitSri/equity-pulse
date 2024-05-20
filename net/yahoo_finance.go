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
)

const (
	yFDefaultScheme = "https"
	yFBaseURL       = "finance.yahoo.com"
	yFQuery1URL     = "query1." + yFBaseURL
	yFQuery2URL     = "query2." + yFBaseURL
	yFMediaURL      = "ncp-gw-finance.media.yahoo.com"
)

type ArticleSet map[string]*model.Article

type YahooFinanceDataProvider struct {
	client *http.Client
}

func NewYahooFinanceDataProvider(c *http.Client) *YahooFinanceDataProvider {
	return &YahooFinanceDataProvider{
		client: c,
	}
}

func (y *YahooFinanceDataProvider) RetrieveStockNews(ticker string) (*model.News, error) {
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

	set := make(ArticleSet)
	for _, match := range matches {
		set[match[1]] = nil
	}

	if err := y.retrieveArticle(set); err != nil {
		return nil, err
	}

	n := make(model.News, 0, len(set))
	for _, v := range set {
		n = append(n, v)
	}

	return &n, nil
}

func (y *YahooFinanceDataProvider) retrieveArticle(s ArticleSet) error {
	p := map[string][]string{
		"device": {"desktop"},
		"site":   {"finance"},
		"lang":   {"en-CA"},
		"region": {"CA"},
		"bot":    {"0"},
	}

	u, err := util.BuildURL(fmt.Sprintf("https://%s", yFBaseURL), "/caas/content/article", p)
	if err != nil {
		return err
	}

	if err := y.verifyYahooFinanceCookie(u); err != nil {
		return err
	}

	h := http.Header{}
	h.Set(util.HostHeader, yFMediaURL)
	h.Set(util.AcceptHeader, util.ContentTypeJSON)

	for uuid := range s {
		fmt.Println("Retrieving article uuid", uuid)
		cr := u.Query()
		cr.Del("uuid")
		cr.Add("uuid", uuid)
		u.RawQuery = cr.Encode()

		ya := &yahoo.Article{}
		if err := util.FetchAndDecode(y.client, u, http.MethodGet, h, ya); err != nil {
			fmt.Printf("unable to fetch %s: %s", u, err)
		}

		if len(ya.Items) == 0 {
			fmt.Println("no data for article uuid", uuid)
			continue
		}

		r := ya.Items[0]
		a := model.NewArticleBuilder().Title(r.Schema.Default.Headline).Summary(r.Schema.Default.Description).Url(r.Schema.Default.MainEntityOfPage).DatePublished(r.Schema.Default.DatePublished).DateUpdated(r.Schema.Default.DateModified).Keywords(r.Schema.Default.Keywords).Build()
		s[uuid] = a
	}

	return nil
}

func (y *YahooFinanceDataProvider) RetrieveCompanyProfile(ticker string) (*model.CompanyProfile, error) {
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
	cp := model.NewCompanyProfileBuilder().Address(r.SummaryProfile.Address1).City(r.SummaryProfile.City).State(r.SummaryProfile.State).Zip(r.SummaryProfile.Zip).Country(r.SummaryProfile.Country).Phone(r.SummaryProfile.Phone).Website(r.SummaryProfile.Website).Industry(r.SummaryProfile.Industry).Sector(r.SummaryProfile.Sector).LongBusinessSummary(r.SummaryProfile.LongBusinessSummary).FullTimeEmployee(uint32(r.SummaryProfile.FullTimeEmployees)).InvestorRelationWebsite(r.SummaryProfile.IrWebsite).Build()

	return cp, nil
}

func (y *YahooFinanceDataProvider) RetrieveStockStatistics(ticker string) (*model.StockStatistics, error) {
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
	s := model.NewStockStatisticsBuilder().FiftyTwoWeekChange(r.FiftyTwoWeekChange.Fmt).Beta(r.Beta.Fmt).BookValue(r.BookValue.Fmt).EnterpriseToEbitda(r.EnterpriseToEbitda.Fmt).EnterpriseToRevenue(r.EnterpriseToRevenue.Fmt).EnterpriseValue(r.EnterpriseValue.Fmt).FiveYearAverageReturn(r.FiveYearAverageReturn.Fmt).FloatShares(r.FloatShares.Fmt).ForwardEps(r.ForwardEps.Fmt).ForwardPE(r.ForwardPE.Fmt).HeldPercentInsiders(r.HeldPercentInsiders.Fmt).HeldPercentInstitutions(r.HeldPercentInstitutions.Fmt).ImpliedSharesOutstanding(r.ImpliedSharesOutstanding.Fmt).LastDividendDate(r.LastDividendDate.Fmt).LastDividendValue(r.LastDividendValue.Fmt).NetIncomeToCommon(r.NetIncomeToCommon.Fmt).PegRatio(r.PegRatio.Fmt).PriceToBook(r.PriceToBook.Fmt).ProfitMargins(r.ProfitMargins.Fmt).RevenueQuarterlyGrowth(r.RevenueQuarterlyGrowth.Fmt).SharesOutstanding(r.SharesOutstanding.Fmt).SharesShort(r.SharesShort.Fmt).ShortRatio(r.ShortRatio.Fmt).TotalAssets(r.TotalAssets.Fmt).TrailingEps(r.TrailingEps.Fmt).Yield(r.Yield.Fmt).YtdReturn(r.YtdReturn.Fmt).Build()
	return s, nil
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

func (y *YahooFinanceDataProvider) RetrieveStockTickerInfo(ticker string) (*model.TickerInfo, error) {
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
	t := model.NewTickerInfoBuilder().Symbol(r.Symbol).Exchange(r.Exchange).QuoteType(r.QuoteType).LongName(r.LongName).ShortName(r.ShortName).Build()

	return t, nil
}

func (y *YahooFinanceDataProvider) getCrumb() (string, error) {
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
