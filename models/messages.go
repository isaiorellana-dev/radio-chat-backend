package models

import "time"

type MessageWithUser struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}
