package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserToReturn struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLogin struct {
	Nickname string
	Pin      string
}

type UserRegister struct {
	ID        int       `json:"id"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

type AppClaims struct {
	UserID   int
	Nickname string
	jwt.StandardClaims
}
