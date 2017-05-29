package models

// Member represents a Meetup group member
type Member struct {
	Name   string  `json:"name"`
	Status string  `json:"status"`
	ID     int     `json:"id"`
	Topics []Topic `json:"topics"`
}

// Members wraps a slice of Member and also contains meta-fields from the meetup API response
type Members struct {
	Members    []Member `json:"results"`
	TotalCount int      `json:"total_count"`
	Count      int      `json:"count"`
}
