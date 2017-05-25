package meetup

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const baseURL = "https://api.meetup.com"

// Clienter
type Clienter interface{}

type ClientOpts struct {
	APIKey string
}

// Client
type Client struct {
	hc   *http.Client
	opts *ClientOpts
}

// NewClient ...
func NewClient() *Client {
	return &Client{
		hc: &http.Client{
			Timeout: time.Duration(time.Second * 20),
		},
	}
}

// Options configures the client with global details that are provided by the consumer
func (c *Client) Options(opts *ClientOpts) {
	c.opts = opts
}

// call
func (c *Client) call(method, uri string, data *bytes.Buffer, result interface{}) error {
	var req *http.Request
	//req.Header.Add("", "")

	var err error

	endpoint := baseURL + uri

	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(method, endpoint, nil)
		if err != nil {
			return err
		}

	case http.MethodPost:
		req, err = http.NewRequest(method, endpoint, data)
		if err != nil {
			return err
		}
	}
	defer func() { req.Close = true }()

	res, err := c.hc.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(result)
}
