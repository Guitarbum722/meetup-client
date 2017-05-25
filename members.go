package meetup

import (
	"net/http"
	"net/url"
	"strconv"
)

// Member represents a Meetup group member
type Member struct {
	Name   string     `json:"name"`
	Status string     `json:"status"`
	ID     int        `json:"id"`
	Topics []Interest `json:"topics"`
}

type Members struct {
	Members    []Member `json:"results"`
	TotalCount int      `json:"total_count"`
	Count      int      `json:"count"`
}

// Interest describes a topic that the member lists as a topic of 'interest' on their profile
type Interest struct {
	Name   string `json:"name"`
	URLKey string `json:"urlkey"`
	ID     int    `json:"id"`
}

// Members returns all of the members that belong to the specified meetup group
func (c *Client) Members(groupID int) (Members, error) {
	var members Members

	query := "?group_id=" + strconv.Itoa(groupID) + "&key=" + c.opts.APIKey
	u, error := url.Parse(query)
	if error != nil {
		return Members{}, error
	}
	q := u.Query()
	u.RawQuery = q.Encode()
	uri := "/2/members?" + u.RawQuery

	err := c.call(http.MethodGet, uri, nil, &members)
	if err != nil {
		return Members{}, err
	}

	return members, nil
}

// Member returns the meetup profile data for a single member
func (c *Client) Member(memberID int) (Member, error) {
	var member Member

	uri := "/2/member/" + strconv.Itoa(memberID) + "?key=" + c.opts.APIKey
	err := c.call(http.MethodGet, uri, nil, &member)
	if err != nil {
		return member, err
	}

	return member, nil
}
