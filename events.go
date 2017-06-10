package meetup

import (
	"net/http"

	"github.com/Guitarbum722/meetup-client/models"
)

const (
	smartRadius        = "smart"
	findEventsEndpoint = "/find/events"
)

// EventsByGeo returns event data based on latitude, longitude and radius respectively.
// Radius can be a value of 'smart', 'global', or in between 0 and 100
// If an empty string is passed for radius, then 'smart' will be used as a default
func (c *Client) EventsByGeo(lat, lon, radius string) ([]models.Event, error) {
	v := c.urlValues()
	if radius == "" {
		radius = smartRadius
	}
	v.Add("radius", radius)
	v.Add("lon", lon)
	v.Add("lat", lat)

	uri := findEventsEndpoint + queryStart + v.Encode()

	var events []models.Event
	if err := c.call(http.MethodGet, uri, nil, &events); err != nil {
		return nil, err
	}

	return events, nil
}
