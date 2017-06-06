package models

// Group describes a meetup group
type Group struct {
	Name        string `json:"name"`
	URLName     string `json:"urlname"`
	ID          int    `json:"id"`
	Link        string `json:"link"`
	MemberCount int    `json:"members"`
	Organizer   Member `json:"organizer"`
}

// Groups wraps a slice of Group for unmarshalling the results array.
// It also contains meta fields from the response.
type Groups struct {
	Groups     []Group `json:"results"`
	TotalCount int     `json:"total_count"`
	Count      int     `json:"count"`
}
