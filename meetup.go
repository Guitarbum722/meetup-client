package meetup

import "bytes"

const baseURL = "https://api.meetup.com"

// Clienter
type Clienter interface{}

// Client
type Client struct{}

// call
func (c *Client) call(method, endpoint string, data *bytes.Buffer, result interface{}) error {
	return nil
}
