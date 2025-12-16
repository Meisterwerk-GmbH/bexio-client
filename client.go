package bexio

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const defaultBaseURL = "https://api.bexio.com/3.0"

// Client holds HTTP configuration for talking to the Bexio V3 API.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	UserAgent  string
}

// Option lets callers override client defaults.
type Option func(*Client)

// WithBaseURL overrides the default API URL (useful for testing).
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.BaseURL = url
	}
}

// WithHTTPClient allows injecting a custom http.Client (timeouts, tracing, etc.).
func WithHTTPClient(cl *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = cl
	}
}

// WithUserAgent sets a custom user agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) {
		c.UserAgent = ua
	}
}

// NewClient constructs a client with sensible defaults.
func NewClient(token string, opts ...Option) *Client {
	c := &Client{
		BaseURL:    defaultBaseURL,
		Token:      token,
		HTTPClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// NewRequest prepares an HTTP request with base URL, bearer auth, and JSON headers.
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	if c == nil {
		return nil, fmt.Errorf("nil client")
	}
	url := strings.TrimSuffix(c.BaseURL, "/") + "/" + strings.TrimPrefix(path, "/")
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	if c.Token == "" {
		return nil, fmt.Errorf("missing API token")
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do executes the request using the configured HTTP client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c == nil {
		return nil, fmt.Errorf("nil client")
	}
	if req == nil {
		return nil, fmt.Errorf("nil request")
	}
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return httpClient.Do(req)
}
