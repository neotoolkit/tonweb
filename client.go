package tonweb

import (
	"net/http"
	"net/url"
)

type Client struct {
	Schema   string
	Host     string
	BasePath string
	client   *http.Client
}

func NewClient() *Client {
	return &Client{
		Schema:   "https",
		Host:     "toncenter.com",
		BasePath: "/api/v2",
		client:   http.DefaultClient,
	}
}

func (c *Client) Accounts() Accounts {
	return Accounts{c}
}

func (c *Client) Blocks() Blocks {
	return Blocks{c}
}

func (c *Client) url() *url.URL {
	var res url.URL
	res.Scheme = c.Schema
	res.Host = c.Host
	return &res
}

type Response struct {
	Ok     bool   `json:"ok"`
	Result any    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Code   int32  `json:"code,omitempty"`
}
