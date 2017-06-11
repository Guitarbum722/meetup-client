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
}

// Venue ...
type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
