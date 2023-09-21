package application

import "auth_api/core/domain"

type DefaultAuthenticationService struct {
	authRepository    domain.AuthRepository
	sessionRepository domain.SessionRepository
}

func (this *DefaultAuthenticationService) New(AuthRepository domain.AuthRepository, SessionRepository domain.SessionRepository) {
	this.authRepository = AuthRepository
	this.sessionRepository = SessionRepository
}

func (this *DefaultAuthenticationService) Authenticate(credentials domain.Credentials) (domain.Session, error) {
	user, error := this.authRepository.SignIn(credentials)
	if error != nil {
		return domain.Session{}, error
	}
	session, error := this.sessionRepository.CreateSession(user)
	return session, error

}
func (this *DefaultAuthenticationService) Logout(session domain.Session) error {
	error := this.sessionRepository.SignOut(session.Id)
	if error != nil {
		return error
	}
	return nil
}
