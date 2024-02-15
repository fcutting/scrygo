package scrygo

import (
	"fmt"
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

func (c Client) sendGet(url string) (responseBody []byte, err error) {
	resp, err := c.HTTPClient.Get(url)

	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
