package models

// Item a
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Duration    int64  `json:"duration"`
	CreatedDate int64  `json:"created_date"`
}
