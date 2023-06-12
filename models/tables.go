package models

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname" validate:"required,min=3,alphaunicode"`
}

type Message struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}
