package domain

type AuthRepository interface {
	SignIn(credentials Credentials) (User, error)
	SignUp(credentials Credentials, user User) (User, error)
}

type SessionRepository interface {
	CreateSession(user User) (Session, error)
	GetSession(sessionID string) (Session, error)
	SignOut(sessionID string) (bool, error)
}
