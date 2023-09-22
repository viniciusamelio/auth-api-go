package domain

import (
	"auth_api/core"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSUTShouldReturnErrorWhenSignInDoesSo(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	authRepo.On("SignIn", mock.Anything).Return(
		User{},
		&core.DefaultError{
			Message: "Invalid credentials",
		},
	)
	sut := AuthenticationService{
		authRepository:    authRepo,
		sessionRepository: sessionRepo,
	}

	_, err := sut.Authenticate(Credentials{
		Username: "test",
		Password: "test",
	})

	assert.Equal(t, "Invalid credentials", err.Error())
}

func TestSUTShouldReturnErrorWhenSignInIsOkButSessionCreationFails(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	authRepo.On("SignIn", mock.Anything).Return(
		User{
			Email: faker.Email(),
			Name:  faker.Name(),
			Id:    faker.UUIDHyphenated(),
		},
		nil,
	)
	sessionRepo.On("CreateSession", mock.Anything).Return(
		Session{}, &core.DefaultError{
			Message: "Session could not be created",
		},
	)
	sut := AuthenticationService{
		authRepository:    authRepo,
		sessionRepository: sessionRepo,
	}

	_, err := sut.Authenticate(Credentials{
		Username: "test",
		Password: "test",
	})

	assert.Equal(t, "Session could not be created", err.Error())
}

func TestSUTShouldReturnSessionWhenSignInIsOk(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	userOutput := User{
		Email: faker.Email(),
		Name:  faker.Name(),
		Id:    faker.UUIDHyphenated(),
	}
	authRepo.On("SignIn", mock.Anything).Return(
		userOutput,
		nil,
	)
	sessionRepo.On("CreateSession", mock.Anything).Return(
		Session{
			User:      userOutput,
			Id:        faker.UUIDHyphenated(),
			Active:    true,
			Token:     faker.UUIDHyphenated(),
			ExpiresAt: time.Now().Add(24 * time.Hour),
		}, nil,
	)
	sut := AuthenticationService{
		authRepository:    authRepo,
		sessionRepository: sessionRepo,
	}

	session, err := sut.Authenticate(Credentials{
		Username: "test",
		Password: "test",
	})

	assert.Equal(t, session.User, userOutput)
	assert.Equal(t, session.Active, true)
	assert.Nil(t, err)
}
