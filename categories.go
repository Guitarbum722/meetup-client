package meetup

import (
	"github.com/Guitarbum722/meetup-client/models"
	"net/http"
)

const categoriesEndpoint = "/2/categories"

// Categories returns available meetup categories
func (c *Client) Categories() (*models.Categories, error) {
	v := c.urlValues()

	uri := categoriesEndpoint + queryStart + v.Encode()

	var categories models.Categories
	if err := c.call(http.MethodGet, uri, nil, &categories); err != nil {
		return nil, err
	}

	return &categories, nil
}
