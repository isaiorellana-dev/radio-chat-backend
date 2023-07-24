package models

import (
	"github.com/golang-jwt/jwt"
)

type UserToReturn struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type UserLogin struct {
	Nickname string
	Pin      string
}

type UserRegister struct {
	ID        int    `json:"id"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"created_at"`
}

type UserLoginDB struct {
	ID       int
	Nickname string
	Pin      string
	Role     string
}

type AppClaims struct {
	UserID   int
	Nickname string
	RolID    int
	jwt.StandardClaims
}
