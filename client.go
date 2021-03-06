package tonweb

import (
	"net/http"
	"net/url"
)

type Client struct {
	Schema   string
	Host     string
	BasePath string
	APIKey   string
	client   *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		Schema:   "https",
		Host:     "toncenter.com",
		BasePath: "/api/v2",
		APIKey:   apiKey,
		client:   http.DefaultClient,
	}
}

func (c *Client) Accounts() Accounts {
	return Accounts{c}
}

func (c *Client) Blocks() Blocks {
	return Blocks{c}
}

func (c *Client) Transactions() Transactions {
	return Transactions{c}
}

func (c *Client) Run() Run {
	return Run{c}
}

func (c *Client) Send() Send {
	return Send{c}
}

func (c *Client) JsonRPC() JsonRPC {
	return JsonRPC{c}
}

func (c *Client) url() *url.URL {
	var res url.URL
	res.Scheme = c.Schema
	res.Host = c.Host
	return &res
}

type Response struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
	Code   int32       `json:"code,omitempty"`
}
