package internal

import "net/http"

// transport is an http.RoundTripper that makes Authorization HTTP requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a token from the supplied Sources.
//
// transport is a low-level mechanism. Most code will use the
// higher-level Config.Client method instead.
type transport struct {
	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// APIKey for use the API
	APIKey string
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	// TODO : add bearer token check
	return t.Base.RoundTrip(r)
}
