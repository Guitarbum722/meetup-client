package meetup

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/Guitarbum722/meetup-client/models"
)

const baseURL = "https://api.meetup.com"

// Clienter
type Clienter interface {
	Members(int) (*models.Members, error)
	Member(int) (*models.Member, error)
	GroupByID([]int) (*models.Groups, error)
	GroupByURLName([]string) (*models.Groups, error)
	GroupByOrganizer([]int) (*models.Groups, error)
	GroupByZip(int) (*models.Groups, error)
	Categories() (*models.Categories, error)
}

type ClientOpts struct {
	APIKey string
}

// Client
type Client struct {
	hc   *http.Client
	opts *ClientOpts
}

// NewClient ...
func NewClient(opts *ClientOpts) Clienter {
	return &Client{
		hc: &http.Client{
			Timeout: time.Duration(time.Second * 20),
		},
		opts: opts,
	}
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

// returns url.Values map with initialized API key
func (c *Client) urlValues() url.Values {
	v := url.Values{}
	v.Set("key", c.opts.APIKey)
	return v
}
