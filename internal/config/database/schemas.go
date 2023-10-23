package database

import (
	"time"

	"gorm.io/gorm"
)

type Schema struct {
	gorm.Model
}

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Hash      string    `json:"hash"`
	Email     string    `json:"email" gorm:"index:unique"`
	Id        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	gorm.Model
	Id        string    `gorm:"primaryKey"`
	Token     string    `json:"token"`
	UserId    string    `json:"user_id"`
	Active    bool      `json:"active"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
