package models

type User struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname" validate:"required,min=3,alphaunicode"`
}
