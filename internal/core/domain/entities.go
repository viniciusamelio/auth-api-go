package domain

import "time"

type Credentials struct {
	Username string
	Password string
}

type User struct {
	Email string
	Name  string
	Id    string
}

type Session struct {
	Id        string
	User      User
	ExpiresAt time.Time
	Active    bool
	Token     string
}

func (self Session) isExpired() bool {
	return self.ExpiresAt.Before(time.Now())
}

func (self *Session) clean() {
	self.User = User{}
	self.ExpiresAt = time.Time{}
	self.Id = ""
	self.Active = false
}
