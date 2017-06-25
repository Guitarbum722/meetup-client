package meetup

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Guitarbum722/meetup-client/models"
)

const (
	v3EventsEndpoint     = "/events"
	v2EventEndpoint      = "/2/event"
	v2EventsEndpoint     = "/2/events"
	eventCommentEndpoint = "/2/event_comment"
	eventsConcierge      = "/2/concierge"
	ratingsEndpoint      = "/2/event_rating"
	smartRadius          = "smart"
)

// Commonly used query param or form field names.  These can be used as the options passed to your eopts func
// and the keys in the map the configures the options.
const (
	CommentID     = "comment_id"
	CommentText   = "comment"
	InReplyTo     = "in_reply_to"
	MemberID      = "member_id"
	GroupID       = "group_id"
	EventID       = "event_id"
	EventName     = "name"
	Rating        = "rating"
	GroupURLName  = "group_urlname"
	Description   = "description"
	PublishStatus = "publish_status"
	EventTime     = "time"
)

// Status types to be used in Event queries
const (
	EventCancelled = "cancelled"
	EventDraft     = "draft"
	EventPast      = "past"
	EventProposed  = "proposed"
	EventSuggested = "suggested"
	EventUpcoming  = "upcoming"
)

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
func (c *Client) EventComments(prep func(map[string][]string, url.Values), o map[string][]string) (*models.Comments, error) {
	v := c.urlValues()
	prep(o, v)

	uri := eventCommentEndpoint + "s" + queryStart + v.Encode()

	var comments models.Comments
	if err := c.call(http.MethodGet, uri, nil, &comments); err != nil {
		return nil, err
	}

	return &comments, nil
}

// EventCommentByID returns a single Comment using the provided comment id.
func (c *Client) EventCommentByID(commentID int) (*models.Comment, error) {
	v := c.urlValues()

	uri := eventCommentEndpoint + fwdSlash + strconv.Itoa(commentID) + queryStart + v.Encode()

	var comment models.Comment
	if err := c.call(http.MethodGet, uri, nil, &comment); err != nil {
		return nil, err
	}

	return &comment, nil
}

// EventRatings returns the ratings for the given eventID
// options o is required to have at least an EventID and an optional MemberID
func (c *Client) EventRatings(prep func(map[string][]string, url.Values), o map[string][]string) (*models.Ratings, error) {
	v := c.urlValues()
	prep(o, v)

	uri := ratingsEndpoint + "s" + queryStart + v.Encode()

	var ratings models.Ratings
	if err := c.call(http.MethodGet, uri, nil, &ratings); err != nil {
		return nil, err
	}

	return &ratings, nil
}

// RateEvent posts the provided rating to the specified eventID
// Use EventID and Rating as options
func (c *Client) RateEvent(prep func(map[string][]string, url.Values), o map[string][]string) (*models.Rating, error) {
	v := c.urlValues()

	uri := ratingsEndpoint + queryStart + v.Encode()

	form := url.Values{}
	prep(o, form)

	data := bytes.NewBufferString(form.Encode())

	var rating models.Rating
	if err := c.call(http.MethodPost, uri, data, &rating); err != nil {
		return nil, err
	}

	return &rating, nil
}

// CommentOnEvent posts a comment to the specified event
// Required event options are EventID and CommentText
// Optionally, use InReplyTo to specify the comment ID in which to reply to
func (c *Client) CommentOnEvent(prep func(map[string][]string, url.Values), o map[string][]string) (*models.Comment, error) {
	v := c.urlValues()

	uri := eventCommentEndpoint + queryStart + v.Encode()

	form := url.Values{}
	prep(o, form)

	data := bytes.NewBufferString(form.Encode())

	var comment models.Comment
	if err := c.call(http.MethodPost, uri, data, &comment); err != nil {
		return nil, err
	}
	return &comment, nil
}

// RemoveEventComment deletes a previously posted comment with the provided commentID
func (c *Client) RemoveEventComment(commentID int) error {
	v := c.urlValues()

	uri := eventCommentEndpoint + fwdSlash + strconv.Itoa(commentID) + queryStart + v.Encode()

	var res interface{}
	if err := c.call(http.MethodDelete, uri, &bytes.Buffer{}, &res); err != nil {
		return err
	}

	return nil
}

// LikeComment uses the specified commentID to 'like' it.
func (c *Client) LikeComment(commentID int) error {
	v := c.urlValues()

	uri := eventCommentEndpoint + "_like" + fwdSlash + strconv.Itoa(commentID) + queryStart + v.Encode()

	var res interface{}
	if err := c.call(http.MethodPost, uri, &bytes.Buffer{}, &res); err != nil {
		return err
	}

	return nil
}

// UnlikeComment will remove a previously posted 'like' on the specified commentID
func (c *Client) UnlikeComment(commentID int) error {
	v := c.urlValues()

	uri := eventCommentEndpoint + "_like" + fwdSlash + strconv.Itoa(commentID) + queryStart + v.Encode()

	var res interface{}
	if err := c.call(http.MethodDelete, uri, &bytes.Buffer{}, &res); err != nil {
		return err
	}

	return nil
}

// CreateEvent posts a new event for the given group
// EventOpts required are GroupID, GroupURLName and EventName (name of the event)
// Optional Event Opts supported by this lib include Description, PublishStatus (organizer only)
// You can set more options with the passed prep func and map parameters to this method (see Meetup API docs for a full list)
func (c *Client) CreateEvent(prep func(map[string][]string, url.Values), o map[string][]string) (*models.Event, error) {
	v := c.urlValues()

	uri := v2EventEndpoint + queryStart + v.Encode()

	form := url.Values{}
	prep(o, form)

	data := bytes.NewBufferString(form.Encode())

	var event models.Event
	if err := c.call(http.MethodPost, uri, data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// UpdateEvent modifies an existing event
func (c *Client) UpdateEvent(eventID string, prep func(map[string][]string, url.Values), o map[string][]string) (*models.Event, error) {
	v := c.urlValues()

	uri := v2EventEndpoint + fwdSlash + eventID + queryStart + v.Encode()

	form := url.Values{}
	prep(o, form)

	data := bytes.NewBufferString(form.Encode())

	var event models.Event
	if err := c.call(http.MethodPost, uri, data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

// DeleteEvent removes a previously posted event with the given eventID
func (c *Client) DeleteEvent(eventID string) error {
	v := c.urlValues()

	uri := v2EventEndpoint + fwdSlash + eventID + queryStart + v.Encode()

	var res interface{}
	if err := c.call(http.MethodDelete, uri, &bytes.Buffer{}, &res); err != nil {
		return err
	}

	return nil
}
