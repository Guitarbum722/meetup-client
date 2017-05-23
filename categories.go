package meetup

// Category
type Category struct {
	Name      string `json:"name"`
	SortName  string `json:"sort_name"`
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
}

// Categories
func (c *Client) Categories() ([]Category, error) {
	return nil, nil
}
