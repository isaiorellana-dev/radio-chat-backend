package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"type:int"`
	Nickname  string    `json:"nickname" validate:"required,min=3,alphaunicode" gorm:"size:20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Message struct {
	ID        int       `json:"id" gorm:"type:int"`
	Body      string    `json:"body" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at"`
	UserID    int       `json:"user_id" gorm:"foreignKey:UserID, type:int"`
}
