package meetup

import (
	"fmt"
	"github.com/Guitarbum722/meetup-client/models"
	"net/http"
	"net/url"
	"strconv"
)

const (
	groupsEndpoint = "/2/groups"
)

// GroupByID returns the data for a single meetup group using the specified groupID
func (c *Client) GroupByID(groupID int) (*models.Groups, error) {
	var groups models.Groups

	v := url.Values{}
	v.Set("group_id", strconv.Itoa(groupID))
	v.Add("key", c.opts.APIKey)

	uri := fmt.Sprintf("%s?%s", groupsEndpoint, v.Encode())

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByURLName returns the data for a single meetup group using the specified urlName
func (c *Client) GroupByURLName(urlName string) (*models.Groups, error) {
	var groups models.Groups

	v := url.Values{}
	v.Set("group_urlname", urlName)
	v.Add("key", c.opts.APIKey)

	uri := groupsEndpoint + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByOrganizer returns the data for a single meetup group using the specified organizerID
func (c *Client) GroupByOrganizer(organizerID int) (*models.Groups, error) {
	var groups models.Groups

	v := url.Values{}
	v.Set("organizer_id", strconv.Itoa(organizerID))
	v.Add("key", c.opts.APIKey)

	uri := groupsEndpoint + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}
