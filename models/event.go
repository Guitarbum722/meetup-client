package models

// Event describes a meetup event and pertinent information such as id, comments, etc.
type Event struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Link     string `json:"link"`
	YesRSVP  int    `json:"yes_rsvp_count"`
	Waitlist int    `json:"waitlist_count"`
	Group    `json:"group"`
	Venue    `json:"venue"`
}

// Comment is a meetup event comment
type Comment struct {
}

// Venue ...
type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
