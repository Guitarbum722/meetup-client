package meetup

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Guitarbum722/meetup-client/models"
)

const (
	v3EventsEndpoint     = "/events"
	v2EventsEndpoint     = "/2/events"
	eventCommentEndpoint = "/2/event_comment"
	eventsConcierge      = "/2/concierge"
	EventCancelled       = "cancelled"
	EventDraft           = "draft"
	EventPast            = "past"
	EventProposed        = "proposed"
	EventSuggested       = "suggested"
	EventUpcoming        = "upcoming"

	smartRadius = "smart"
)

// EventOpts represents different query params that can be used for Event related endpoints.
type EventOpts struct {
	CommentIDs []string
	MemberIDs  []string
	GroupIDs   []string
	EventIDs   []string
}

// EventsByGeo returns event data based on latitude, longitude and radius respectively.
// Radius can be a value of 'smart', or in between 0.5 and 100
// If an empty string is passed for radius, then 'smart' will be used as a default
func (c *Client) EventsByGeo(lat, lon, radius string) (*models.Events, error) {
	v := c.urlValues()
	if radius == "" {
		radius = smartRadius
	}
	v.Add("radius", radius)
	v.Add("lon", lon)
	v.Add("lat", lat)

	uri := eventsConcierge + queryStart + v.Encode()

	var events models.Events
	if err := c.call(http.MethodGet, uri, nil, &events); err != nil {
		return nil, err
	}

	return &events, nil
}

// EventsByGroup returns event data for the specified group with its urlName.
// Use these contstants to input status:
// EventCancelled, EventDraft, EventPast, EventProposed, EventSuggested, EventUpcoming
func (c *Client) EventsByGroup(urlName string, status []string, desc bool) (*models.Events, error) {
	v := c.urlValues()
	v.Add("status", strings.Join(status, ","))
	v.Add("group_urlname", urlName)
	v.Add("desc", strconv.FormatBool(desc))

	uri := v2EventsEndpoint + queryStart + v.Encode()

	var events models.Events
	if err := c.call(http.MethodGet, uri, nil, &events); err != nil {
		return nil, err
	}

	return &events, nil
}

// EventsByGroupID returns event data for the specified groupID.
// Use these contstants to input status:
// EventCancelled, EventDraft, EventPast, EventProposed, EventSuggested, EventUpcoming
func (c *Client) EventsByGroupID(groupID int, status []string, desc bool) (*models.Events, error) {
	v := c.urlValues()
	v.Add("status", strings.Join(status, ","))
	v.Add("group_id", strconv.Itoa(groupID))
	v.Add("desc", strconv.FormatBool(desc))

	uri := v2EventsEndpoint + queryStart + v.Encode()

	var events models.Events
	if err := c.call(http.MethodGet, uri, nil, &events); err != nil {
		return nil, err
	}

	return &events, nil
}

// EventByID returns a single event with the specified group's url name and event ID
func (c *Client) EventByID(urlName, eventID string) (*models.Event, error) {
	v := c.urlValues()

	uri := fwdSlash + urlName + v3EventsEndpoint + fwdSlash + eventID + queryStart + v.Encode()

	var event models.Event
	if err := c.call(http.MethodGet, uri, nil, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

// EventComments returns comments based on the query criteria provided with the EventOpts
func (c *Client) EventComments(eOpts *EventOpts) (*models.Comments, error) {
	v := c.urlValues()
	v = eventOptsToURLVals(eOpts, v)

	uri := eventCommentEndpoint + "s" + queryStart + v.Encode()

	var comments models.Comments
	if err := c.call(http.MethodGet, uri, nil, &comments); err != nil {
		return nil, err
	}

	return &comments, nil
}

// helper that prepares params with EventOpts
func eventOptsToURLVals(e *EventOpts, v url.Values) url.Values {
	if len(e.CommentIDs) > 0 {
		v.Add("comment_id", strings.Join(e.CommentIDs, ","))
	}
	if len(e.EventIDs) > 0 {
		v.Add("event_id", strings.Join(e.EventIDs, ","))
	}
	if len(e.GroupIDs) > 0 {
		v.Add("group_id", strings.Join(e.GroupIDs, ","))

	}
	if len(e.MemberIDs) > 0 {
		v.Add("member_id", strings.Join(e.MemberIDs, ","))

	}

	return v
}
