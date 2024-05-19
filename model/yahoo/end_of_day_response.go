package yahoo

import (
	"encoding/json"
	"time"
)

type EndOfDay struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Error  Error    `json:"error"`
	Result []Result `json:"result"`
}

type Indicators struct {
	Adjclose []AdjClose `json:"adjclose"`
	Quote    []Quote    `json:"quote"`
}

type AdjClose struct {
	Close []float64 `json:"adjclose"`
}

type Quote struct {
	Close  []float64 `json:"close"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Volume []int64   `json:"volume"`
}

type Meta struct {
	ChartPreviousClose   float64              `json:"chartPreviousClose"`
	Currency             string               `json:"currency"`
	CurrentTradingPeriod CurrentTradingPeriod `json:"currentTradingPeriod"`
	DataGranularity      string               `json:"dataGranularity"`
	ExchangeName         string               `json:"exchangeName"`
	ExchangeTimezoneName string               `json:"exchangeTimezoneName"`
	FiftyTwoWeekHigh     int64                `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64              `json:"fiftyTwoWeekLow"`
	FirstTradeDate       int64                `json:"firstTradeDate"`
	FullExchangeName     string               `json:"fullExchangeName"`
	Gmtoffset            int64                `json:"gmtoffset"`
	HasPrePostMarketData bool                 `json:"hasPrePostMarketData"`
	InstrumentType       string               `json:"instrumentType"`
	PriceHint            int64                `json:"priceHint"`
	Range                string               `json:"range"`
	RegularMarketDayHigh int64                `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64              `json:"regularMarketDayLow"`
	RegularMarketPrice   float64              `json:"regularMarketPrice"`
	RegularMarketTime    int64                `json:"regularMarketTime"`
	RegularMarketVolume  int64                `json:"regularMarketVolume"`
	Symbol               string               `json:"symbol"`
	Timezone             string               `json:"timezone"`
	ValidRanges          []string             `json:"validRanges"`
}

type CurrentTradingPeriod struct {
	Post    TradingPeriod `json:"post"`
	Pre     TradingPeriod `json:"pre"`
	Regular TradingPeriod `json:"regular"`
}

type TradingPeriod struct {
	End       int64  `json:"end"`
	Gmtoffset int64  `json:"gmtoffset"`
	Start     int64  `json:"start"`
	Timezone  string `json:"timezone"`
}

type UnixTime struct {
	time.Time
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}

	t.Time = time.Unix(timestamp, 0)
	return nil
}
