package models

// Event describes a meetup event and pertinent information such as id, comments, etc.
type Event struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	YesRSVP  int    `json:"yes_rsvp_count"`
	Waitlist int    `json:"waitlist_count"`
	Group    Group  `json:"group"`
	Venue    Venue  `json:"venue"`
}

// Events wraps a slice of Event for unmarshalling the results array.
// It also contains meta fields from the response.
type Events struct {
	Events     []Event `json:"results"`
	TotalCount int     `json:"total_count"`
	Count      int     `json:"count"`
}

// Comment is a meetup event comment
type Comment struct {
	MemberID    int    `json:"member_id"`
	MemberName  string `json:"member_name"`
	CommentID   int    `json:"event_comment_id"`
	EventID     string `json:"event_id"`
	GroupID     int    `json:"group_id"`
	CommentText string `json:"comment"`
}

// Comments wraps a slice of Event for unmarshalling the results array.
// It also contains meta fields from the response.
type Comments struct {
	Comments   []Comment `json:"results"`
	TotalCount int       `json:"total_count"`
	Count      int       `json:"count"`
}

// Venue represents the location for a meetup
type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Rating is submitted for a particular meetup event by a member
type Rating struct {
	MemberID   int    `json:"member_id"`
	MemberName string `json:"member_name"`
	EventID    string `json:"event_id"`
	GroupID    int    `json:"group_id"`
	Rating     int    `json:"rating"`
}

// Ratings wraps a slice of Rating for unmarshalling the results array.
// It also contains meta fields from the response.
type Ratings struct {
	Ratings    []Rating `json:"results"`
	TotalCount int      `json:"total_count"`
	Count      int      `json:"count"`
}
