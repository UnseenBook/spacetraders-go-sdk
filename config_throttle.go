package stc

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

type ThrottledTransport struct {
	rt http.RoundTripper
	rl *rate.Limiter
}

func (tt *ThrottledTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	err := tt.rl.Wait(r.Context())
	if err != nil {
		if r.Body != nil {
			_ = r.Body.Close()
		}
		return nil, err
	}

	return tt.rt.RoundTrip(r)
}

func (c *Configuration) AddThrottler(period time.Duration, burst int) error {
	if c.HTTPClient == nil {
		return reportError("to configure a throttler first add a HTTPClient to the Configuration")
	}
	if c.HTTPClient.Transport == nil {
		return reportError("to configure a throttler first add a Transport to the HTTPClient")
	}
	tt := &ThrottledTransport{
		rt: c.HTTPClient.Transport,
		rl: rate.NewLimiter(rate.Every(period), burst),
	}
	c.HTTPClient.Transport = tt

	return nil
}
