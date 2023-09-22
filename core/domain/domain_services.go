package domain

import (
	"auth_api/core"
)

// Authentication domain service
type AuthenticationService struct {
	authRepository    AuthRepository
	sessionRepository SessionRepository
}

func (this *AuthenticationService) New(AuthRepository AuthRepository, SessionRepository SessionRepository) {
	this.authRepository = AuthRepository
	this.sessionRepository = SessionRepository
}

func (this *AuthenticationService) Authenticate(Credentials Credentials) (Session, error) {
	user, error := this.authRepository.SignIn(Credentials)
	if error != nil {
		return Session{}, error
	}
	session, error := this.sessionRepository.CreateSession(user)
	return session, error

}
func (this *AuthenticationService) Logout(Session Session) error {
	error := this.sessionRepository.SignOut(Session.Id)
	if error != nil {
		return error
	}
	return nil
}

func (this *AuthenticationService) SignUp(Credentials Credentials, User User) (User, error) {
	user, error := this.authRepository.SignUp(Credentials, User)
	if error != nil {
		return User, error
	}
	return user, nil
}

// Session domain service
type SessionService struct {
	sessionRepository SessionRepository
}

func (this *SessionService) New(SessionRepository SessionRepository) {
	this.sessionRepository = SessionRepository
}

func (this *SessionService) RecoverSession(SessionID string) (Session, error) {
	session, error := this.sessionRepository.GetSession(SessionID)
	if error != nil {
		return Session{}, error
	}
	if session.isExpired() && session.Active {
		error := this.sessionRepository.SignOut(session.Id)
		if error != nil {
			return Session{}, core.DefaultError{
				Message: "Session could not be expired",
			}
		}
		return Session{}, core.DefaultError{
			Message: "Invalid session",
		}
	}
	return session, nil
}
