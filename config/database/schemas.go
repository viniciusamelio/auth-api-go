package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Hash      string
	Email     string
	Id        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	Id        string `gorm:"primaryKey"`
	Token     string
	UserId    string
	Active    bool
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
