package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"type:int;primaryKey;not null;autoIncrement;unique"`
	Nickname  string    `json:"nickname" validate:"required,min=3,alphaunicode" gorm:"size:20;unique;not null"`
	Pin       string    `json:"pin" validate:"required,min=4,numeric" gorm:"type:string;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	RolID     int       `gorm:"type:int;not null;default:1"`
	Role      Role      `json:"Role" gorm:"foreignKey:RolID"`
}

type Message struct {
	ID        int       `json:"id" gorm:"type:int;primaryKey;not null;autoIncrement;unique"`
	Body      string    `json:"body" validate:"required,min=1" gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UserID    int       `json:"user_id" gorm:"foreignKey:UserID"`
}

type Role struct {
	ID          int          `json:"id" gorm:"type:int;primaryKey;not null;autoIncrement;unique"`
	Name        string       `json:"name" gorm:"not null;size:40;unique"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	ID   int    `json:"id" gorm:"type:int;primaryKey;not null;autoIncrement;unique"`
	Name string `json:"name" gorm:"not null;size:40;unique"`
}
