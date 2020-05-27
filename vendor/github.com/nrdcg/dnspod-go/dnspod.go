// Package dnspod implements a client for the dnspod API.
//
// In order to use this package you will need a dnspod account and your API Token.
package dnspod

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	libraryVersion   = "0.4"
	defaultBaseURL   = "https://dnsapi.cn/"
	defaultUserAgent = "dnspod-go/" + libraryVersion

	// apiVersion       = "v1"
	defaultTimeout   = 5
	defaultKeepAlive = 30
)

// dnspod API docs: https://www.dnspod.cn/docs/info.html

// CommonParams is the commons parameters.
type CommonParams struct {
	LoginToken   string
	Format       string
	Lang         string
	ErrorOnEmpty string
	UserID       string

	Timeout   int
	KeepAlive int
}

func (c CommonParams) toPayLoad() url.Values {
	p := url.Values{}

	if c.LoginToken != "" {
		p.Set("login_token", c.LoginToken)
	}
	if c.Format != "" {
		p.Set("format", c.Format)
	}
	if c.Lang != "" {
		p.Set("lang", c.Lang)
	}
	if c.ErrorOnEmpty != "" {
		p.Set("error_on_empty", c.ErrorOnEmpty)
	}
	if c.UserID != "" {
		p.Set("user_id", c.UserID)
	}

	return p
}

// Status is the status representation.
type Status struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type service struct {
	client *Client
}

// Client is the DNSPod client.
type Client struct {
	// HTTP client used to communicate with the API.
	HTTPClient *http.Client

	// CommonParams used communicating with the dnspod API.
	CommonParams CommonParams

	// Base URL for API requests.
	// Defaults to the public dnspod API, but can be set to a different endpoint (e.g. the sandbox).
	// BaseURL should always be specified with a trailing slash.
	BaseURL string

	// User agent used when communicating with the dnspod API.
	UserAgent string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the dnspod API.
	Domains *DomainsService
	Records *RecordsService
}

// NewClient returns a new dnspod API client.
func NewClient(params CommonParams) *Client {
	timeout := defaultTimeout
	if params.Timeout != 0 {
		timeout = params.Timeout
	}

	keepalive := defaultKeepAlive
	if params.KeepAlive != 0 {
		keepalive = params.KeepAlive
	}

	httpClient := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(timeout) * time.Second,
				KeepAlive: time.Duration(keepalive) * time.Second,
			}).DialContext,
		},
	}

	client := &Client{HTTPClient: &httpClient, CommonParams: params, BaseURL: defaultBaseURL, UserAgent: defaultUserAgent}

	client.common.client = client
	client.Domains = (*DomainsService)(&client.common)
	client.Records = (*RecordsService)(&client.common)

	return client
}

// NewRequest creates an API request.
// The path is expected to be a relative path and will be resolved
// according to the BaseURL of the Client. Paths should always be specified without a preceding slash.
func (c *Client) NewRequest(method, path string, payload url.Values) (*http.Request, error) {
	uri := c.BaseURL + path

	req, err := http.NewRequest(method, uri, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) post(path string, payload url.Values, v interface{}) (*Response, error) {
	return c.Do(http.MethodPost, path, payload, v)
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to decode it.
func (c *Client) Do(method, path string, payload url.Values, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path, payload)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	response := &Response{Response: res}
	err = CheckResponse(res)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, res.Body)
		} else {
			err = json.NewDecoder(res.Body).Decode(v)
		}
	}

	return response, err
}

// A Response represents an API response.
type Response struct {
	*http.Response
}

// An ErrorResponse represents an error caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // human-readable message
}

// Error implements the error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if the status code is different than 2xx. Specific requests
// may have additional requirements, but this is sufficient in most of the cases.
func CheckResponse(r *http.Response) error {
	if code := r.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	err := json.NewDecoder(r.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

// Date custom type.
type Date struct {
	time.Time
}

// UnmarshalJSON handles the deserialization of the custom Date type.
func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("date should be a string, got %s: %w", data, err)
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date: %w", err)
	}

	d.Time = t

	return nil
}
