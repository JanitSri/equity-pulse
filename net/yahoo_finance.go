package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/util"
	"golang.org/x/net/publicsuffix"
)

const (
	baseURL1 = "https://query1.finance.yahoo.com"
	baseURL2 = "https://query2.finance.yahoo.com"
)

type YahooFinanceDataProvider struct {
	client *http.Client
	cache  cache
}

func NewYahooFinanceDataProvider(c *http.Client) *YahooFinanceDataProvider {
	// TODO: setup cache
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

	u, _ := util.BuildURL(baseURL1, fmt.Sprintf("/v1/finance/quoteType/%s", ticker), p)

	if err := y.verifyCookie(u); err != nil {
		return nil, err
	}

	var q *model.QuoteInfoYahooFinance
	err := y.fetchAndDecode(u, http.MethodGet, nil, &q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if q == nil || len(q.QuoteType.Result) == 0 {
		return nil, fmt.Errorf("empty result for QuoteType")
	}

	r := q.QuoteType.Result[0]
	t := model.NewTickerInfoBuilder().Symbol(r.Symbol).Exchange(r.Exchange).QuoteType(r.QuoteType).LongName(r.LongName).ShortName(r.ShortName).Build()

	return t, nil
}

func (y *YahooFinanceDataProvider) fetchAndDecode(u *url.URL, method string, h http.Header, target interface{}) error {
	req := y.buildRequest(u, method, h)

	res, err := y.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned non-200 status: %d", res.Request.URL.String(), res.StatusCode)
	}

	defer res.Body.Close()

	d := json.NewDecoder(res.Body)
	err = d.Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func (y *YahooFinanceDataProvider) buildRequest(u *url.URL, m string, nH http.Header) *http.Request {
	h := map[string][]string{
		userAgentHeader: {"Mozilla/5.0 (compatible; MyBot/1.0)"},
		hostHeader:      {"query1.finance.yahoo.com"},
		acceptHeader:    {"application/json"},
	}

	for k, v := range nH {
		h[k] = v
	}

	return &http.Request{
		URL:    u,
		Header: h,
		Method: m,
	}
}

func (y *YahooFinanceDataProvider) verifyCookie(iU *url.URL) error {
	cs := y.client.Jar.Cookies(iU)
	e := len(cs) == 0
	for _, cookie := range cs {
		e = util.IsCookieExpired(cookie)
		if !e {
			break
		}
	}

	if e {
		u, err := url.Parse("https://finance.yahoo.com/")
		if err != nil {
			return err
		}

		h := map[string][]string{
			acceptHeader: {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8"},
		}
		r := y.buildRequest(u, http.MethodGet, h)
		if _, err := y.client.Do(r); err != nil {
			return err
		}

		// check if cookie is valid for url
		if len(y.client.Jar.Cookies(iU)) == 0 {
			return fmt.Errorf("cookie is not valid for %s", iU)
		}
	}

	return nil
}

func (y *YahooFinanceDataProvider) getCrumb() (string, error) {
	// TODO: get crumb query parameter
	return "", nil
}
