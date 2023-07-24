package models

type MessageWithUser struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"created_at"`
}
