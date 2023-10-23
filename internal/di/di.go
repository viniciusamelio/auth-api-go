package dependencies

import (
	config "auth_api/internal/config"
	application "auth_api/internal/core/application"
	domain "auth_api/internal/core/domain"
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
