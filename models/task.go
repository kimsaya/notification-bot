package models

// Task a
type Task struct {
	UserID      string `json:"user_id"`
	ItemID      string `json:"item_id"`
	TaskTitle   string `json:"task_title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedDate string `json:"creater_date"`
}
