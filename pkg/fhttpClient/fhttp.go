package fhttp

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

type Client struct {
}

var timeout = time.Millisecond * 5000

func NewClient() *Client {
	return &Client{}
}

func (h *Client) Request(method string, url string, body []byte, queryParams map[string]string, headers map[string]string) (responseBody []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(method)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	var queryCount int
	for k, v := range queryParams {
		switch queryCount {
		case 0:
			url += fmt.Sprintf("?%s=%s", k, v)
		default:
			url += fmt.Sprintf("&%s=%s", k, v)
		}
		queryCount++
	}

	req.SetBody(body)
	req.SetRequestURI(url)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	client := &fasthttp.Client{}

	if err = client.DoTimeout(req, res, timeout); err != nil {
		return
	}

	statusCode = res.StatusCode()

	responseBody = make([]byte, len(res.Body()))
	copy(responseBody, res.Body())
	req.SetConnectionClose()
	return
}
