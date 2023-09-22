package domain

// Authentication domain service
type AuthenticationService struct {
	authRepository    AuthRepository
	sessionRepository SessionRepository
}

func (this *AuthenticationService) New(AuthRepository AuthRepository, SessionRepository SessionRepository) {
	this.authRepository = AuthRepository
	this.sessionRepository = SessionRepository
}

func (this *AuthenticationService) Authenticate(credentials Credentials) (Session, error) {
	user, error := this.authRepository.SignIn(credentials)
	if error != nil {
		return Session{}, error
	}
	session, error := this.sessionRepository.CreateSession(user)
	return session, error

}
func (this *AuthenticationService) Logout(session Session) error {
	error := this.sessionRepository.SignOut(session.Id)
	if error != nil {
		return error
	}
	return nil
}

// Session domain service
type SessionService struct {
	sessionRepository SessionRepository
}

func (this *SessionService) RecoverSession(sessionID string) (Session, error) {
	session, error := this.sessionRepository.GetSession(sessionID)
	return session, error
}
