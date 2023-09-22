package application

import (
	"auth_api/config/database"
	"auth_api/core"
	domain "auth_api/core/domain"
	"time"

	"github.com/google/uuid"
)

type DefaultSessionRepository struct {
	DatabaseDatasource database.GoOrmDatabase
}

func (this *DefaultSessionRepository) CreateSession(user domain.User) (domain.Session, error) {
	sessionDto := database.Session{
		UserId:    user.Id,
		Active:    true,
		Token:     uuid.NewString(),
		Id:        uuid.NewString(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := this.DatabaseDatasource.Save(&sessionDto)

	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Session could not be created")
		return domain.Session{}, defaultError
	}

	return domain.Session{
		Id:        sessionDto.Id,
		User:      user,
		ExpiresAt: sessionDto.ExpiresAt,
		Active:    sessionDto.Active,
		Token:     sessionDto.Token,
	}, nil
}
func (this *DefaultSessionRepository) GetSession(sessionID string) (domain.Session, error) {
	var sessionDto database.Session
	var userDto database.User
	result := this.DatabaseDatasource.Select("*").Where("id = ?", sessionID).First(&sessionDto)
	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Invalid session")
		return domain.Session{}, defaultError
	}
	userQueryResult := this.DatabaseDatasource.Select("*").Where("id = ?", sessionDto.UserId).First(&userDto)
	if userQueryResult.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Invalid session")
		return domain.Session{}, defaultError
	}
	return domain.Session{
		Id: sessionDto.Id,
		User: domain.User{
			Id:    userDto.Id,
			Name:  userDto.Name,
			Email: userDto.Email,
		},
		ExpiresAt: sessionDto.ExpiresAt,
		Active:    sessionDto.Active,
		Token:     sessionDto.Token,
	}, nil
}
func (this *DefaultSessionRepository) SignOut(sessionID string) error {
	var sessionDto database.Session
	foundSession := this.DatabaseDatasource.Select("*").Where("id = ?", sessionID).First(sessionDto)
	if foundSession.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Invalid session")
		return defaultError
	}

	sessionDto.Active = false
	sessionDto.ExpiresAt = time.Now()
	sessionDto.UpdatedAt = time.Now()

	result := this.DatabaseDatasource.Save(&sessionDto)

	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Session could not be expired")
		return defaultError
	}

	return nil
}
