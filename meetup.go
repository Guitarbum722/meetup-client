package meetup

import (
	"bytes"
	"encoding/json"
	"github.com/Guitarbum722/meetup-client/models"
	"net/http"
	"time"
)

const baseURL = "https://api.meetup.com"

// Clienter
type Clienter interface {
	Members(groupID int) (*models.Members, error)
	Member(memberID int) (*models.Member, error)
	GroupByID(groupID int) (*models.Groups, error)
	GroupByURLName(urlName string) (*models.Group, error)
	GroupByOrganizer(organizerID int) (*models.Group, error)
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
