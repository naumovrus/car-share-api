package verifypassapi

import (
	"car-api/config"
	"net/http"
	"net/url"
)

const (
	methodPost = "POST"
)

type client struct {
	httpClient http.Client
	config     *config.Config
}

func New(httpClient http.Client, config *config.Config) API {
	return &client{httpClient: httpClient, config: config}
}

func (c *client) authHeaders() map[string][]string {
	var header http.Header
	header.Set("Content-Type", "json")
	return header
}

func (c *client) SendRequest() {
	var (
		req = http.Request{
			Method: methodPost,
			Header: c.authHeaders(),
			URL:    &url.URL{Host: c.config.DocsVerifyService.Url, Path: "api/verify/passport"},
		}
	)
	c.httpClient.Do(&req)
}
