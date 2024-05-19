package model

type Interval string

const (
	Interval1m      Interval = "1 minute"
	Interval2m      Interval = "2 minutes"
	Interval5m      Interval = "5 minutes"
	Interval15m     Interval = "15 minutes"
	Interval30m     Interval = "30 minutes"
	Interval60m     Interval = "60 minutes"
	Interval90m     Interval = "90 minutes"
	Interval1h      Interval = "1 hour"
	Interval1d      Interval = "1 day"
	Interval5d      Interval = "5 days"
	Interval1wk     Interval = "1 week"
	Interval1mo     Interval = "1 month"
	Interval3mo     Interval = "3 months"
	InvalidInterval Interval = "invalid interval"
)

type IntervalConverter struct {
	interval Interval
}

func NewIntervalConverter(interval Interval) *IntervalConverter {
	return &IntervalConverter{interval}
}

func (ic *IntervalConverter) ConvertToYahooFinanceInterval() Interval {
	switch ic.interval {
	case Interval1m:
		return "1m"
	case Interval2m:
		return "2m"
	case Interval5m:
		return "5m"
	case Interval15m:
		return "15m"
	case Interval30m:
		return "30m"
	case Interval60m:
		return "60m"
	case Interval90m:
		return "90m"
	case Interval1h:
		return "1h"
	case Interval1d:
		return "1d"
	case Interval5d:
		return "5d"
	case Interval1wk:
		return "1wk"
	case Interval1mo:
		return "1mo"
	case Interval3mo:
		return "3mo"
	default:
		return InvalidInterval
	}
}
