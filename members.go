package meetup

import (
	"net/http"
	"strconv"

	"github.com/Guitarbum722/meetup-client/models"
)

const (
	memberEndpoint = "/2/member"
	queryStart     = "?"
	fwdSlash       = "/"
)

// Members returns all of the members that belong to the specified meetup group
func (c *Client) Members(groupID int) (*models.Members, error) {
	var members models.Members

	v := c.urlValues()
	v.Add("group_id", strconv.Itoa(groupID))

	// append 's' for /members
	uri := memberEndpoint + "s" + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &members); err != nil {
		return nil, err
	}

	return &members, nil
}

// Member returns the meetup profile data for a single member
func (c *Client) Member(memberID int) (*models.Member, error) {
	var member models.Member

	v := c.urlValues()

	uri := memberEndpoint + fwdSlash + strconv.Itoa(memberID) + queryStart + v.Encode()

	if err := c.call(http.MethodGet, uri, nil, &member); err != nil {
		return nil, err
	}

	return &member, nil
}
