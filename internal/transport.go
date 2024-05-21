package internal

import "net/http"

// Transport is an http.RoundTripper that makes Authorization HTTP requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a token from the supplied Sources.
//
// Transport is a low-level mechanism. Most code will use the
// higher-level Config.Client method instead.
type Transport struct {
	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper
}
