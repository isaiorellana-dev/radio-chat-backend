package models

type MessageWithUser struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
	// UserID   int    `json:"user_id"`
	Nickname string `json:"nickname"`
}
