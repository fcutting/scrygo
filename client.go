package scrygo

import (
	"net/http"
	"time"
)

const (
	defaultBaseURL = "https://api.scryfall.com"
	defaultTimeout = 5 * time.Second
)

// The client struct for managing requests to the Scryfall API.
type Client struct {
	// the base URL for requests to the Scryfall API
	BaseURL string

	// request timeout in seconds
	Timeout int

	// the http client used for requests
	HTTPClient *http.Client
}
