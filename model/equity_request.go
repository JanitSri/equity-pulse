package model

type EquityRequest struct {
	ticker string
	name   string
}

func (e *EquityRequest) Ticker() string {
	return e.ticker
}

func (e *EquityRequest) Name() string {
	return e.name
}

type EquityRequestBuilder struct {
	ticker string
	name   string
}

func (e *EquityRequestBuilder) Ticker(t string) *EquityRequestBuilder {
	e.ticker = t
	return e
}

func (e *EquityRequestBuilder) Name(n string) *EquityRequestBuilder {
	e.name = n
	return e
}

func (e *EquityRequestBuilder) Build() *EquityRequest {
	return &EquityRequest{
		ticker: e.ticker,
		name:   e.name,
	}
}

func NewEquityRequestBuilder() *EquityRequestBuilder {
	return &EquityRequestBuilder{}
}
