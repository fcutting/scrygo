package scrygo

import (
	"io"
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

// Sends a GET request to the provided URL and returns the response body while
// handling the closing of the response body.
func (c Client) sendGet(url string) (responseBody []byte, err error) {
	resp, err := c.HTTPClient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
