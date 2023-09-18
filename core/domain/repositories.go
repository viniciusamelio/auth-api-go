package domain

type AuthRepository interface {
	SignIn(credentials Credentials) (User, error)
	SignUp(credentials Credentials, user User) (User, error)
}

type SessionRepository interface {
	CreateSession(user User) (Session, error)
	GetSession(sessionID Uuid) (Session, error)
	SignOut(sessionID Uuid) error
}
