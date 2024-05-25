package util

import (
	"time"

	"go.uber.org/zap"
)

const (
	TimeZoneUTC          = "UTC"
	TimeZoneNewYork      = "America/New_York"
	TimeZoneLosAngeles   = "America/Los_Angeles"
	TimeZoneChicago      = "America/Chicago"
	TimeZoneDenver       = "America/Denver"
	TimeZoneLondon       = "Europe/London"
	TimeZoneParis        = "Europe/Paris"
	TimeZoneBerlin       = "Europe/Berlin"
	TimeZoneMoscow       = "Europe/Moscow"
	TimeZoneBeijing      = "Asia/Shanghai"
	TimeZoneTokyo        = "Asia/Tokyo"
	TimeZoneSydney       = "Australia/Sydney"
	TimeZoneAuckland     = "Pacific/Auckland"
	TimeZoneHongKong     = "Asia/Hong_Kong"
	TimeZoneSingapore    = "Asia/Singapore"
	TimeZoneMumbai       = "Asia/Kolkata"
	TimeZoneDubai        = "Asia/Dubai"
	TimeZoneSaoPaulo     = "America/Sao_Paulo"
	TimeZoneJohannesburg = "Africa/Johannesburg"
	TimeZoneCairo        = "Africa/Cairo"
)

// TODO: write unit tests
func FormatDate(date time.Time, timeZone, layout string) string {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		zap.L().Sugar().Errorf("could not load time zone: %s", err)
		loc, _ = time.LoadLocation(TimeZoneUTC)
		zap.L().Sugar().Infof("using default time zone: %s", loc)
	}

	updatedTimeZone := date.In(loc)
	output := updatedTimeZone.Format(layout)
	if output == layout {
		zap.L().Sugar().Errorf("could not use layout: %s", layout)
		output = updatedTimeZone.Format(time.RFC1123)
		zap.L().Sugar().Info("using default layout: RFC1123")
	}

	return output
}
