package models

// Topic describes a topic of interest that is listed for a member, group, etc.
type Topic struct {
	Name   string `json:"name"`
	URLKey string `json:"urlkey"`
	ID     int    `json:"id"`
}
