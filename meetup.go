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
	EventsByGeo(string, string, string) (*models.Events, error)
	EventsByGroup(string, []string, bool) (*models.Events, error)
	EventByID(string, string) (*models.Event, error)
	EventsByGroupID(int, []string, bool) (*models.Events, error)
	EventComments(func(map[string][]string, url.Values), map[string][]string) (*models.Comments, error)
	EventCommentByID(int) (*models.Comment, error)
	EventRatings(func(map[string][]string, url.Values), map[string][]string) (*models.Ratings, error)
	RateEvent(func(map[string][]string, url.Values), map[string][]string) (*models.Rating, error)
	CommentOnEvent(func(map[string][]string, url.Values), map[string][]string) (*models.Comment, error)
	LikeComment(int) error
	UnlikeComment(int) error
	RemoveEventComment(int) error
	CreateEvent(func(map[string][]string, url.Values), map[string][]string) (*models.Event, error)
	UpdateEvent(string, func(map[string][]string, url.Values), map[string][]string) (*models.Event, error)
	DeleteEvent(string) error
}

// ClientOpts contains options to be passed in when creating a new
// meetup client value
type ClientOpts struct {
	APIKey     string
	HTTPClient *http.Client
}

// Client represents a meetup client
type Client struct {
	hc   *http.Client
	opts *ClientOpts
}

// NewClient creates a new Meetup client with the given parameters
func NewClient(opts *ClientOpts) Clienter {
	if opts.HTTPClient != nil {
		return &Client{
			hc:   opts.HTTPClient,
			opts: opts,
		}
	}
	return &Client{
		hc: &http.Client{
			Timeout: time.Duration(time.Second * 20),
		},
		opts: opts,
	}
}

// call acts as a broker for the packages HTTP requests to the meetup.com API
func (c *Client) call(method, uri string, data *bytes.Buffer, result interface{}) error {
	var req *http.Request

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
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			return err
		}
	case http.MethodDelete:
		req, err = http.NewRequest(method, endpoint, nil)
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
