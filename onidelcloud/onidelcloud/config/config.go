package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

type Config struct {
	APIKey      string
	APIEndpoint string
}

func (c *Config) Client() (*Client, error) {
	client := &Client{
		HTTPClient: &http.Client{},
		BaseURL:    c.APIEndpoint,
		APIKey:     c.APIKey,
	}
	log.Printf("[INFO] Onidel Cloud Client configured for URL: %s", client.BaseURL)
	return client, nil
}

func (c *Client) doRequest(method, path string, queryParams map[string]string, body, result interface{}) error {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("error parsing URL: %s", err)
	}

	q := u.Query()
	for k, v := range queryParams {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error marshaling request body: %s", err)
		}
		reqBody = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, u.String(), reqBody)
	if err != nil {
		return fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %s", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("error unmarshaling response: %s", err)
		}
	}

	return nil
}

func (c *Client) Get(path string, queryParams map[string]string, result interface{}) error {
	return c.doRequest("GET", path, queryParams, nil, result)
}

func (c *Client) Post(path string, body, result interface{}) error {
	return c.doRequest("POST", path, nil, body, result)
}

func (c *Client) PostWithQuery(path string, queryParams map[string]string, body, result interface{}) error {
	return c.doRequest("POST", path, queryParams, body, result)
}

func (c *Client) Patch(path string, body, result interface{}) error {
	return c.doRequest("PATCH", path, nil, body, result)
}

func (c *Client) Put(path string, body, result interface{}) error {
	return c.doRequest("PUT", path, nil, body, result)
}

func (c *Client) Delete(path string, queryParams map[string]string) error {
	return c.doRequest("DELETE", path, queryParams, nil, nil)
}

func (c *Client) DeleteWithResult(path string, queryParams map[string]string, result interface{}) error {
	return c.doRequest("DELETE", path, queryParams, nil, result)
}
