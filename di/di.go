package dependencies

import (
	"auth_api/config"
	"auth_api/core/application"
	"auth_api/core/domain"
)

func NewAuthRepository() domain.AuthRepository {
	return &application.DefaultAuthRepository{
		Database: config.Database,
	}
}

func NewSessionRepository() domain.SessionRepository {
	return &application.DefaultSessionRepository{
		DatabaseDatasource: config.Database,
	}
}

func NewAuthenticationService() application.ApplicationAuthenticationService {
	instance := &application.AuthenticationService{}
	domainAuthService := &domain.AuthenticationService{}
	domainAuthService.New(NewAuthRepository(), NewSessionRepository())
	instance.New(
		domainAuthService,
	)
	return instance
}

func NewSessionService() application.ApplicationSessionService {
	instance := &application.SessionService{}
	domainSessionService := domain.SessionService{}
	domainSessionService.New(NewSessionRepository())
	instance.New(
		domainSessionService,
	)
	return instance
}
