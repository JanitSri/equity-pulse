package yahoo

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type QuoteSummary struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}

type Image struct {
	Type   string `json:"@type"`
	Tag    string `json:"tag"`
	Height int64  `json:"height"`
	URL    string `json:"url"`
	Width  int64  `json:"width"`
}

type Provider struct {
	DisplayName string `json:"displayName"`
	Type        string `json:"@type"`
	Logo        Logo   `json:"logo"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type Result struct {
	SummaryProfile            SummaryProfile       `json:"summaryProfile"`
	DefaultKeyStatistics      DefaultKeyStatistics `json:"defaultKeyStatistics"`
	FinancialData             FinancialData        `json:"financialData"`
	Exchange                  string               `json:"exchange"`
	ExchangeTimezoneName      string               `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string               `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds     string               `json:"gmtOffSetMilliseconds"`
	IsEsgPopulated            bool                 `json:"isEsgPopulated"`
	LongName                  string               `json:"longName"`
	Market                    string               `json:"market"`
	MessageBoardID            string               `json:"messageBoardId"`
	QuoteType                 string               `json:"quoteType"`
	ShortName                 string               `json:"shortName"`
	Symbol                    string               `json:"symbol"`
}
