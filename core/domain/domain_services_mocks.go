package domain

import "github.com/stretchr/testify/mock"

type mockAuthRepository struct {
	mock.Mock
}

func (this *mockAuthRepository) SignIn(credentials Credentials) (User, error) {
	args := this.Called(credentials)
	return args.Get(0).(User), args.Error(1)
}
func (this *mockAuthRepository) SignUp(credentials Credentials, user User) (User, error) {
	args := this.Called(credentials, user)
	return args.Get(0).(User), args.Error(1)
}

type mockSessionRepository struct {
	mock.Mock
}

func (this *mockSessionRepository) CreateSession(user User) (Session, error) {
	args := this.Called(user)
	return args.Get(0).(Session), args.Error(1)
}
func (this *mockSessionRepository) GetSession(sessionID string) (Session, error) {
	args := this.Called(sessionID)
	return args.Get(0).(Session), args.Error(1)
}
func (this *mockSessionRepository) SignOut(sessionID string) (bool, error) {
	args := this.Called(sessionID)
	return args.Get(0).(bool), args.Get(1).(error)
}
