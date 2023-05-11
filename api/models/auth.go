package models

type User struct {
	ID        string `gorm:"primaryKey"`
	Firstname string `json:"firstname" gorm:"not_null"`
	Lastname  string `json:"lastname" gorm:"not_null"`
	Username  string `json:"username" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not_null"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RedisType struct {
	UserID        string `json:"userid"`
	Authenticated bool   `json:"authenticated"`
}

type SessionAuthdata struct {
	SessionId string `json:"session_id"`
	Validity  int    `json:"valid_till"`
}
