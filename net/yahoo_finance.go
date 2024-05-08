package net

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/util"
)

const baseURL1 string = "https://query1.finance.yahoo.com"
const baseURL2 string = "https://query2.finance.yahoo.com"

type YahooFinanceDataProvider struct {
	client *http.Client
	cache  cache
}

func NewYahooFinanceDataProvider(client *http.Client) *YahooFinanceDataProvider {
	// TODO: setup cache
	return &YahooFinanceDataProvider{
		client: client,
	}
}

// todo: handle errors -> return response and error and let upstream handle it

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

	var q model.QuoteInfoYahooFinance

	err := y.fetchAndDecode(u, http.MethodGet, &q)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r := q.QuoteType.Result[0]
	t := model.NewTickerInfoBuilder().Symbol(r.Symbol).Exchange(r.Exchange).QuoteType(r.QuoteType).LongName(r.LongName).ShortName(r.ShortName).Build()

	return t, nil
}

func (y *YahooFinanceDataProvider) fetchAndDecode(u *url.URL, method string, target interface{}) error {
	req := y.buildRequest(u, method)

	res, err := y.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s returned non-200 status: %d", res.Request.URL.String(), res.StatusCode)
	}

	defer res.Body.Close()

	// when the response is not needed -- for retrieving the cookie
	if target != nil {
		d := json.NewDecoder(res.Body)
		err = d.Decode(&target)
		if err != nil {
			return err
		}
	}

	return nil
}

func (y *YahooFinanceDataProvider) buildRequest(u *url.URL, m string) *http.Request {
	h := map[string][]string{
		"User-Agent": {"Mozilla/5.0 (compatible; MyBot/1.0)"},
		"Host":       {"query1.finance.yahoo.com"},
		"Accept":     {"application/json"},
	}

	return &http.Request{
		URL:    u,
		Header: h,
		Method: m,
	}
}

func (y *YahooFinanceDataProvider) addCookie() {
	// TODO: add logic to verify cookie and update
}
