// Package solarmanager implements a client for the SolarManager API.
// See https://external-web.solar-manager.ch.
package solarmanager

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const (
	defaultBaseURL = "https://cloud.solar-manager.ch/"
	userAgent      = "go-solarmanager"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Verbose   bool
	Username  string
	Password  string

	client *http.Client
}

// NewClient returns a new SolarManager API client using the supplied credentials.
// If a nil httpClient is provided, a new http.Client will be used.
func NewClient(httpClient *http.Client, baseURL *url.URL, username, password string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	if baseURL == nil {
		baseURL, _ = url.Parse(defaultBaseURL)
	}

	return &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
		Username:  username,
		Password:  password,
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.Verbose {
		if d, err := httputil.DumpRequest(req, true); err == nil {
			log.Println(string(d))
		}
	}

	resp, err := c.client.Do(req)

	if c.Verbose {
		if d, err := httputil.DumpResponse(resp, true); err == nil {
			log.Println(string(d))
		}
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}
