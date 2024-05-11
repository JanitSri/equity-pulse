package yahoo

type QuoteInfo struct {
	QuoteType QuoteType `json:"quoteType"`
}

type QuoteType struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}
