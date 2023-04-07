package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	host       string
	httpClient *http.Client
}

func NewClient(host string, timeout time.Duration) Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return Client{
		host:       host,
		httpClient: client,
	}
}

func (c *Client) Do(
	method, endpoint string, headers map[string]string, params map[string]string,
	body []byte,
) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	requestBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, baseURL, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}
