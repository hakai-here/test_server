package models

type TodoActivity struct {
	ID      string `gorm:"primaryKey" json:"id"`
	UserID  string
	Todo    string `gorm:"not_null" json:"todo"`
	DueDate string `gorm:"not_null" json:"due-date"`
	Action  string `gorm:"not_null" json:"action"`
}
