package models

// Task a
type Task struct {
	ID          string `json:"id"` //UserID-BookID-Index-Activity
	Name        string `json:"name"`
	UserID      string `json:"user_id"`
	ItemID      string `json:"item_id"`
	Status      string `json:"status"`
	LastUpdate  int64  `json:"last_update"`
	CreatedDate int64  `json:"creater_date"`
}

type TaskDTO struct {
	Name         string
	IsTranslated bool
	IsEdited     bool
	IsPosted     bool
}
