package models

// User a
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	LastActive string `json:"last_active"`
	JointDate  string `json:"joint_date"`
}
