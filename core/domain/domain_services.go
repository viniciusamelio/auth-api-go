package domain

import "auth_api/core"

type DomainAuthenticationService interface {
	Authenticate(Credentials Credentials) (Session, error)
	SignUp(Credentials Credentials, User User) (User, error)
	Logout(Session Session) (bool, error)
}

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
func (this *AuthenticationService) Logout(Session Session) (bool, error) {
	_, error := this.sessionRepository.SignOut(Session.Id)
	if error != nil {
		return false, error
	}
	return true, nil
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
		success, _ := this.sessionRepository.SignOut(session.Id)
		if !success {
			return Session{}, core.DefaultError{
				Message: "Session could not be expired",
			}
		}
		return Session{}, core.DefaultError{
			Message: "Invalid session",
		}
	} else if session.isExpired() {
		return Session{}, core.DefaultError{
			Message: "Invalid session",
		}
	}
	return session, nil
}
