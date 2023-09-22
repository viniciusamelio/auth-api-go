package core

import "time"

type CredentialsDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDto struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Id    string `json:"id"`
}

type SessionDto struct {
	Id        string    `json:"id"`
	User      UserDto   `json:"user"`
	ExpiresAt time.Time `json:"expires_at"`
	Active    bool      `json:"active"`
	Token     string    `json:"token"`
}
