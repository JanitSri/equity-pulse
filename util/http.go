package util

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	UserAgentHeader   = "User-Agent"
	AcceptHeader      = "Accept"
	HostHeader        = "Host"
	ContentTypeHeader = "Content-Type"

	ContentTypeJSON = "application/json"
	ContentTypeText = "text/plain"
	UserAgent       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
)

type TextBodyResponse struct {
	Body string `json:"body"`
}

func BuildURL(base, path string, params url.Values) (*url.URL, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		zap.L().Sugar().Errorf("could not parse %s: %s", base, err)
		return nil, err
	}

	baseURL.Path += path

	baseURL.RawQuery = params.Encode()

	return baseURL, nil
}

func BuildRequest(u *url.URL, method string, nH http.Header) *http.Request {
	h := http.Header{}
	h.Set(UserAgentHeader, UserAgent)

	for k := range nH {
		h.Set(k, nH.Get(k))
	}

	return &http.Request{
		URL:    u,
		Header: h,
		Method: method,
	}
}

func FetchAndDecode(c *http.Client, u *url.URL, method string, h http.Header, target interface{}) error {
	req := BuildRequest(u, method, h)

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// TODO: implement back off for status 429 - Too Many Requests
	if res.StatusCode != http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s returned non-200 status: %d body: %s", res.Request.URL.String(), res.StatusCode, string(b))
	}

	// target is nil when the decoding is not needed - i.e. retrieving cookie
	var dErr error
	if target != nil {
		if strings.Contains(strings.ToLower(res.Header.Get(ContentTypeHeader)), ContentTypeText) {
			dErr = DecodeTextContentType(res.Body, target)
		} else {
			dErr = DecodeJSONContentType(res.Body, target)
		}
	}

	return dErr
}

func VerifyCookie(c *http.Client, targetUrl, cookieUrl *url.URL, h http.Header) error {
	cs := c.Jar.Cookies(targetUrl)
	e := len(cs) == 0
	for _, cookie := range cs {
		e = IsCookieExpired(cookie)
		if !e {
			break
		}
	}

	if e {
		if err := FetchAndDecode(c, cookieUrl, http.MethodGet, h, nil); err != nil {
			return nil
		}

		// check if cookie is valid for url
		if len(c.Jar.Cookies(targetUrl)) == 0 {
			return fmt.Errorf("cookie is not valid for %s", targetUrl)
		}
	}

	return nil
}

func IsCookieExpired(cookie *http.Cookie) bool {
	return time.Now().After(cookie.Expires)
}
