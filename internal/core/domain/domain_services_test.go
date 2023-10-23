package domain

import (
	"auth_api/internal/core"
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

func TestSUTShouldReturnErrorWhenSignOutFails(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	sessionRepo.On("SignOut", mock.Anything).Return(
		false, &core.DefaultError{
			Message: "Session could not be signed out",
		},
	)
	sut := AuthenticationService{
		sessionRepository: sessionRepo,
		authRepository:    authRepo,
	}

	_, error := sut.Logout(Session{
		Id: faker.UUIDHyphenated(),
	})

	assert.NotNil(t, error)
	assert.Equal(t, "Session could not be signed out", error.Error())
}

func TestSUTShouldReturnErrorWhenAuthRepoSignUpFails(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	authRepo.On("SignUp", mock.Anything, mock.Anything).Return(
		User{}, &core.DefaultError{
			Message: "User could not be created",
		},
	)
	sut := AuthenticationService{
		sessionRepository: sessionRepo,
		authRepository:    authRepo,
	}

	_, error := sut.SignUp(Credentials{
		Username: "test",
		Password: "test",
	}, User{
		Email: faker.Email(),
		Name:  faker.Name(),
		Id:    faker.UUIDHyphenated(),
	})

	assert.NotNil(t, error)
	assert.Equal(t, "User could not be created", error.Error())
}
func TestSUTShouldReturnUserWhenAuthRepoSignSucceeds(t *testing.T) {
	authRepo := new(mockAuthRepository)
	sessionRepo := new(mockSessionRepository)
	authRepo.On("SignUp", mock.Anything, mock.Anything).Return(
		User{
			Email: faker.Email(),
			Name:  faker.Name(),
			Id:    faker.UUIDHyphenated(),
		}, nil,
	)
	sut := AuthenticationService{
		sessionRepository: sessionRepo,
		authRepository:    authRepo,
	}

	user, _ := sut.SignUp(Credentials{
		Username: "test",
		Password: "test",
	}, User{
		Email: faker.Email(),
		Name:  faker.Name(),
		Id:    faker.UUIDHyphenated(),
	})

	assert.IsType(t, User{}, user)
	assert.NotNil(t, user.Email)
	assert.NotNil(t, user.Name)
	assert.NotNil(t, user.Id)
}

func TestSUTShouldReturnErrorWhenSessionRepoDoesSo(t *testing.T) {
	sessionRepo := new(mockSessionRepository)
	sessionRepo.On("GetSession", mock.Anything).Return(
		Session{}, &core.DefaultError{
			Message: "Session could not be created",
		},
	)
	sut := SessionService{
		sessionRepository: sessionRepo,
	}

	_, error := sut.RecoverSession(faker.UUIDHyphenated())

	assert.NotNil(t, error)
	assert.Equal(t, "Session could not be created", error.Error())
}

func TestSUTShouldReturnErrorWhenSessionRepoReturnsASessionButItIsExpiredAndSignOutFails(t *testing.T) {
	sessionRepo := new(mockSessionRepository)
	sessionId := faker.UUIDHyphenated()
	sessionRepo.On("GetSession", mock.Anything).Return(
		Session{
			Id:        sessionId,
			Active:    true,
			Token:     faker.UUIDHyphenated(),
			ExpiresAt: time.Now().AddDate(-1, 0, 0),
		}, nil,
	)
	sessionRepo.On("SignOut", sessionId).Return(
		false, core.DefaultError{
			Message: "Invalid session",
		},
	)
	sut := SessionService{
		sessionRepository: sessionRepo,
	}

	_, error := sut.RecoverSession(faker.UUIDHyphenated())

	assert.NotNil(t, error)
	assert.Equal(t, "Session could not be expired", error.Error())
}

func TestSUTShouldReturnErrorWhenSessionRepoReturnsASessionWhichIsActiveButExpired(t *testing.T) {
	sessionRepo := new(mockSessionRepository)
	sessionId := faker.UUIDHyphenated()
	sessionRepo.On("GetSession", mock.Anything).Return(
		Session{
			Id:        sessionId,
			Active:    true,
			Token:     faker.UUIDHyphenated(),
			ExpiresAt: time.Now().AddDate(-1, 0, 0),
		}, nil,
	)
	sessionRepo.On("SignOut", sessionId).Return(
		true, core.DefaultError{},
	)
	sut := SessionService{
		sessionRepository: sessionRepo,
	}

	_, error := sut.RecoverSession(faker.UUIDHyphenated())

	assert.NotNil(t, error)
	assert.Equal(t, "Invalid session", error.Error())
}

func TestSUTShouldReturnErrorWhenSessionRepoReturnsAnExpiredSession(t *testing.T) {
	sessionRepo := new(mockSessionRepository)
	sessionId := faker.UUIDHyphenated()
	sessionRepo.On("GetSession", mock.Anything).Return(
		Session{
			Id:        sessionId,
			Active:    false,
			Token:     faker.UUIDHyphenated(),
			ExpiresAt: time.Now().AddDate(-1, 0, 0),
		}, nil,
	)
	sessionRepo.On("SignOut", sessionId).Return(
		true, core.DefaultError{},
	)
	sut := SessionService{
		sessionRepository: sessionRepo,
	}

	_, error := sut.RecoverSession(faker.UUIDHyphenated())

	assert.NotNil(t, error)
	assert.Equal(t, "Invalid session", error.Error())
}

func TestSUTShouldReturnSessionWhenItFindsAnActiveOne(t *testing.T) {
	sessionRepo := new(mockSessionRepository)
	sessionId := faker.UUIDHyphenated()
	sessionRepo.On("GetSession", mock.Anything).Return(
		Session{
			Id:        sessionId,
			Active:    true,
			Token:     faker.UUIDHyphenated(),
			ExpiresAt: time.Now().AddDate(0, 0, 1),
		}, nil,
	)
	sut := SessionService{
		sessionRepository: sessionRepo,
	}

	session, _ := sut.RecoverSession(faker.UUIDHyphenated())

	assert.Equal(t, sessionId, session.Id)
	assert.Equal(t, true, session.Active)
	assert.Greater(t, session.ExpiresAt, time.Now())
}
