package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

const (
	defaultBaseURL = "https://api.todoist.com/api/v1"
)

// Client is the Todoist API client.
type Client struct {
	token      string
	baseURL    string
	httpClient *http.Client
}

// Option configures the Client.
type Option func(*Client)

// WithHTTPClient sets a custom http.Client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithBaseURL overrides the default API base URL.
func WithBaseURL(u string) Option {
	return func(c *Client) {
		c.baseURL = u
	}
}

// New creates a new Todoist API client.
func New(token string, opts ...Option) *Client {
	c := &Client{
		token:      token,
		baseURL:    defaultBaseURL,
		httpClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// APIError represents an error returned by the Todoist API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("todoist: API error %d: %s", e.StatusCode, e.Body)
}

func (c *Client) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	u := c.baseURL + path

	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("todoist: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("todoist: read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(data),
		}
	}

	if v != nil && resp.StatusCode != http.StatusNoContent && len(data) > 0 {
		if err := json.Unmarshal(data, v); err != nil {
			return fmt.Errorf("todoist: unmarshal response: %w", err)
		}
	}

	return nil
}

// get performs a GET request.
func (c *Client) get(ctx context.Context, path string, query any, v any) error {
	if query != nil {
		q := structToQuery(query)
		if encoded := q.Encode(); encoded != "" {
			path += "?" + encoded
		}
	}

	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, v)
}

// getList performs a GET request that returns a paginated list, collecting all pages.
func getList[T any](c *Client, ctx context.Context, path string, query any) ([]T, error) {
	var all []T
	cursor := ""

	for {
		q := url.Values{}
		if query != nil {
			q = structToQuery(query)
		}
		if cursor != "" {
			q.Set("cursor", cursor)
		}

		fullPath := path
		if encoded := q.Encode(); encoded != "" {
			fullPath += "?" + encoded
		}

		req, err := c.newRequest(ctx, http.MethodGet, fullPath, nil)
		if err != nil {
			return nil, err
		}

		var page PaginatedResponse[T]
		if err := c.do(req, &page); err != nil {
			return nil, err
		}

		all = append(all, page.Results...)

		if page.NextCursor == nil || *page.NextCursor == "" {
			break
		}
		cursor = *page.NextCursor
	}

	return all, nil
}

// post performs a POST request with a JSON body.
func (c *Client) post(ctx context.Context, path string, body any, v any) error {
	req, err := c.newRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return err
	}
	return c.do(req, v)
}

// delete performs a DELETE request.
func (c *Client) delete(ctx context.Context, path string) error {
	req, err := c.newRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	return c.do(req, nil)
}

// structToQuery converts a struct with `url` tags to url.Values.
func structToQuery(v any) url.Values {
	vals := url.Values{}
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return vals
	}
	rt := rv.Type()
	for i := range rt.NumField() {
		field := rt.Field(i)
		tag := field.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue
		}
		// strip ,omitempty
		name := tag
		if idx := bytes.IndexByte([]byte(tag), ','); idx != -1 {
			name = tag[:idx]
		}
		fv := rv.Field(i)
		if fv.IsZero() {
			continue
		}
		vals.Set(name, fmt.Sprintf("%v", fv.Interface()))
	}
	return vals
}
