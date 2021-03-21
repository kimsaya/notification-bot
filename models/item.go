package models

// Item a
type Item struct {
	ID           string `json:"id"` //Discord's RoleID
	Name         string `json:"name"`
	TranslatorID string `json:"translater_id"`
	EditorID     string `json:"editor_id"`
	PostorID     string `json:"postor_id"`
	Duration     int64  `json:"duration"`
	LastUpdate   int64  `json:"last_update"`
	CreatedDate  int64  `json:"created_date"`
}
