// Package client provides thin HTTP wrappers for weather-api and races-api.
// Each method makes one request and returns the parsed response body,
// the raw status code, and any transport-level error.
//
// Keeping client code separate from test assertions means tests stay readable
// while all the request-building plumbing lives here.
package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Base is a minimal HTTP client shared by all API clients.
type Base struct {
	http       *http.Client
	baseURL    string
	internalKey string
}

func newBase(baseURL, internalKey string, timeout time.Duration) *Base {
	return &Base{
		http:        &http.Client{Timeout: timeout},
		baseURL:     baseURL,
		internalKey: internalKey,
	}
}

// Response holds the raw result of an HTTP call.
type Response struct {
	Status int
	Body   []byte
}

// do executes an HTTP request and returns the response.
func (b *Base) do(method, path string, body any) (*Response, error) {
	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, b.baseURL+path, reqBody)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if b.internalKey != "" {
		req.Header.Set("X-Internal-Key", b.internalKey)
	}

	resp, err := b.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	return &Response{Status: resp.StatusCode, Body: respBody}, nil
}

// Get sends a GET request.
func (b *Base) Get(path string) (*Response, error) {
	return b.do(http.MethodGet, path, nil)
}

// Post sends a POST request with optional JSON body.
func (b *Base) Post(path string, body any) (*Response, error) {
	return b.do(http.MethodPost, path, body)
}

// Parse decodes the response body into v.
func (r *Response) Parse(v any) error {
	if err := json.Unmarshal(r.Body, v); err != nil {
		return fmt.Errorf("parse response (status %d): %w — body: %s", r.Status, err, r.Body)
	}
	return nil
}
