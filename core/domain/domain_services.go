package domain

type AuthenticationService interface {
	Authenticate(credentials Credentials) (Session, error)
	Logout(session Session) error
}

type SessionService interface {
	RecoverSession(sessionID Uuid) (Session, error)
}
