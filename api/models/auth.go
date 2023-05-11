package models

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
