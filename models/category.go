package models

// Category
type Category struct {
	Name      string `json:"name"`
	SortName  string `json:"sort_name"`
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
}

// Categories
type Categories struct {
	Categories []Category `json:"results"`
	TotalCount int        `json:"total_count"`
	Count      int        `json:"count"`
}
