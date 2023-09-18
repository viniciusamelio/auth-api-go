package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Hash  string
	Email string
	Id    string `gorm:"primaryKey"`
}

type Session struct {
	Id        string `gorm:"primaryKey"`
	Token     string
	UserId    string
	Active    bool
	ExpiresAt string
}
