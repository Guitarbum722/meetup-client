package meetup

import (
	"fmt"
	"github.com/Guitarbum722/meetup-client/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	groupsEndpoint = "/2/groups"
)

// GroupByID returns the data for a single meetup group using the specified groupID
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByID(groupIDs ...int) (*models.Groups, error) {
	var groups models.Groups
	var convIDs []string

	for _, id := range groupIDs {
		convIDs = append(convIDs, strconv.Itoa(id))
	}

	v := url.Values{}
	v.Set("key", c.opts.APIKey)
	v.Add("group_id", strings.Join(convIDs, ","))

	uri := fmt.Sprintf("%s?%s", groupsEndpoint, v.Encode())

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByURLName returns the data for a single meetup group using the specified urlName
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByURLName(urlNames ...string) (*models.Groups, error) {
	var groups models.Groups

	v := url.Values{}
	v.Set("key", c.opts.APIKey)
	v.Add("group_urlname", strings.Join(urlNames, ","))

	uri := groupsEndpoint + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByOrganizer returns the data for a single meetup group using the specified organizerID
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByOrganizer(organizerIDs ...int) (*models.Groups, error) {
	var groups models.Groups
	var convIDs []string

	for _, id := range organizerIDs {
		convIDs = append(convIDs, strconv.Itoa(id))
	}

	v := url.Values{}
	v.Set("key", c.opts.APIKey)
	v.Add("organizer_id", strings.Join(convIDs, ","))

	uri := groupsEndpoint + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}
