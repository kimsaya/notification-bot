package models

// User a
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	LastActive int64  `json:"last_active"`
	JointDate  int64  `json:"joint_date"`
}
