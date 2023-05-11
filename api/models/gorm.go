package models

import "gorm.io/gorm"

type User struct {
	ID        string `gorm:"primaryKey"`
	Firstname string `json:"firstname" gorm:"not_null"`
	Lastname  string `json:"lastname" gorm:"not_null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not_null"`
}

type TodoActivity struct {
	gorm.Model
	UserID  string `gorm:"not_null"`
	Todo    string `gorm:"not_null" json:"todo"`
	DueDate string `gorm:"not_null" json:"due-date"`
	Action  string `gorm:"not_null" json:"action"`
}
