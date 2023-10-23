package application

import (
	"auth_api/internal/core"
	domain "auth_api/internal/core/domain"
)

type ApplicationSessionService interface {
	GetSession(SessionId string) (core.SessionDto, error)
}

type SessionService struct {
	sessionService domain.SessionService
}

func (this *SessionService) New(sessionService domain.SessionService) {
	this.sessionService = sessionService
}

func (this *SessionService) GetSession(SessionId string) (core.SessionDto, error) {
	session, error := this.sessionService.RecoverSession(SessionId)
	if error != nil {
		return core.SessionDto{}, error
	}

	return core.SessionDto{
		Id: session.Id,
		User: core.UserDto{
			Email: session.User.Email,
			Name:  session.User.Name,
			Id:    session.User.Id,
		},
		ExpiresAt: session.ExpiresAt,
		Active:    session.Active,
		Token:     session.Token,
	}, nil
}
