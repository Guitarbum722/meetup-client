package meetup

import (
	"fmt"
	"github.com/briandowns/meetup-client/models"
	"net/http"
	"net/url"
	"strconv"
)

const (
	memberEndpoint  = "/2/member"
	membersEndpoint = "/2/members"
)

// Member represents a Meetup group member
type Member struct {
	Name   string         `json:"name"`
	Status string         `json:"status"`
	ID     int            `json:"id"`
	Topics []models.Topic `json:"topics"`
}

// Members wraps a slice of Member and also contains meta-fields from the meetup API response
type Members struct {
	Members    []Member `json:"results"`
	TotalCount int      `json:"total_count"`
	Count      int      `json:"count"`
}

// Members returns all of the members that belong to the specified meetup group
func (c *Client) Members(groupID int) (*Members, error) {
	var members Members

	v := url.Values{}
	v.Set("group_id", strconv.Itoa(groupID))
	v.Set("key", c.opts.APIKey)

	uri := fmt.Sprintf("%s?%s", membersEndpoint, v.Encode())

	if err := c.call(http.MethodGet, uri, nil, &members); err != nil {
		return nil, err
	}

	return &members, nil
}

// Member returns the meetup profile data for a single member
func (c *Client) Member(memberID int) (*Member, error) {
	var member Member

	v := url.Values{}
	v.Set("key", c.opts.APIKey)

	uri := fmt.Sprintf("%s/%d?%s", memberEndpoint, memberID, v.Encode())

	if err := c.call(http.MethodGet, uri, nil, &member); err != nil {
		return nil, err
	}

	return &member, nil
}
