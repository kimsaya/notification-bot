package models

// Item a
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	SubItems    *[]Item `json:"sub_items"`
	Duration    int64   `json:"duration"`
	LastUpdate  int64   `json:"last_update"`
	CreatedDate int64   `json:"created_date"`
}
