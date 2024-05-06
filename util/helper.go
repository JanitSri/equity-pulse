package util

import (
	"fmt"
	"time"
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
func FormatDate(date time.Time, timeZone string) string {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Println("Error loading time zone:", err)
		loc, _ = time.LoadLocation(TimeZoneUTC)
		fmt.Println("Using default time zone:", loc)
	}

	updatedTimeZone := date.In(loc)
	return updatedTimeZone.Format(time.RFC1123)
}
