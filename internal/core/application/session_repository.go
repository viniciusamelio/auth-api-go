package application

import (
	"auth_api/internal/config/database"
	"auth_api/internal/core"
	domain "auth_api/internal/core/domain"
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
	result := this.DatabaseDatasource.Where("id = ?", sessionID).First(&sessionDto)
	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Invalid session")
		return domain.Session{}, defaultError
	}
	userQueryResult := this.DatabaseDatasource.Where("id = ?", sessionDto.UserId).First(&userDto)
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
func (this *DefaultSessionRepository) SignOut(sessionID string) (bool, error) {
	var sessionDto database.Session
	foundSession := this.DatabaseDatasource.Select("*").Where("id = ?", sessionID).First(&sessionDto)
	if foundSession.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Invalid session")
		return false, defaultError
	}

	result := this.DatabaseDatasource.Save(database.Session{
		Id:        sessionDto.Id,
		UserId:    sessionDto.UserId,
		Active:    false,
		Token:     sessionDto.Token,
		ExpiresAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedAt: sessionDto.CreatedAt,
	})

	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("Session could not be expired")
		return false, defaultError
	}

	return true, core.DefaultError{}
}
