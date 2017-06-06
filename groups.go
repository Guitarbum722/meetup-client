package meetup

import (
	"github.com/Guitarbum722/meetup-client/models"
	"net/http"
	"strconv"
	"strings"
)

const (
	groupsEndpoint = "/2/groups"
)

// GroupByID returns the meetup groups using the specified groupIDs
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByID(groupIDs []int) (*models.Groups, error) {
	convIDs := make([]string, 0, len(groupIDs))

	for _, id := range groupIDs {
		convIDs = append(convIDs, strconv.Itoa(id))
	}

	v := c.urlValues()
	v.Add("group_id", strings.Join(convIDs, ","))

	uri := groupsEndpoint + queryStart + v.Encode()

	var groups models.Groups
	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByURLName returns the group data using the specified urlNames
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByURLName(urlNames []string) (*models.Groups, error) {

	v := c.urlValues()
	v.Add("group_urlname", strings.Join(urlNames, ","))

	uri := groupsEndpoint + queryStart + v.Encode()

	var groups models.Groups
	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByOrganizer returns the data the specified meetup groups using organizerIDs
// The response contains an array of results, even if there is only one because the request can
// consist of comma separated values as the group_id parameter.
func (c *Client) GroupByOrganizer(organizerIDs []int) (*models.Groups, error) {
	convIDs := make([]string, 0, len(organizerIDs))

	for _, id := range organizerIDs {
		convIDs = append(convIDs, strconv.Itoa(id))
	}

	v := c.urlValues()
	v.Add("organizer_id", strings.Join(convIDs, ","))

	uri := groupsEndpoint + queryStart + v.Encode()

	var groups models.Groups
	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}

// GroupByZip returns the group data for the specified zip code
func (c *Client) GroupByZip(zipCode int) (*models.Groups, error) {

	v := c.urlValues()
	v.Add("zip", strconv.Itoa(zipCode))

	uri := groupsEndpoint + queryStart + v.Encode()

	var groups models.Groups
	if err := c.call(http.MethodGet, uri, nil, &groups); err != nil {
		return nil, err
	}

	return &groups, nil
}
