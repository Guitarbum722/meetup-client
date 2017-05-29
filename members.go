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

// Members returns all of the members that belong to the specified meetup group
func (c *Client) Members(groupID int) (*models.Members, error) {
	var members models.Members

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
func (c *Client) Member(memberID int) (*models.Member, error) {
	var member models.Member

	v := url.Values{}
	v.Set("key", c.opts.APIKey)

	uri := fmt.Sprintf("%s/%d?%s", memberEndpoint, memberID, v.Encode())

	if err := c.call(http.MethodGet, uri, nil, &member); err != nil {
		return nil, err
	}

	return &member, nil
}
