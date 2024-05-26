package util

import (
	"encoding/json"
	"io"
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

func DecodeTextContentType(r io.Reader, target interface{}) error {
	switch v := target.(type) {
	case *TextBodyResponse:
		b, err := io.ReadAll(r)
		if err != nil {
			return err
		}
		v.Body = string(b)
	default:
		zap.L().Sugar().Error("did not match any content type for decoding")
	}

	return nil
}

func DecodeJSONContentType(r io.Reader, target interface{}) error {
	d := json.NewDecoder(r)
	err := d.Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func ToJSONString(i interface{}) (string, error) {
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
