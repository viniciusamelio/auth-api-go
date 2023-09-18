package application

import (
	domain "auth_api/core/domain"
)

type DefaultAuthRepository struct {
	DatabaseDatasource DatabaseDatasource
}

func (this DefaultAuthRepository) SignIn(credentials domain.Credentials) (domain.User, error) {

	result := this.DatabaseDatasource.Save("users", credentials)
	if result.Failed {
		return domain.User{}, result.Errors[0]
	}
	user := result.Data[0].(map[string]string)
	uuid := domain.Uuid{}
	uuid.Set(user["id"])

	return domain.User{
		Id:    uuid,
		Name:  user["name"],
		Email: user["email"],
	}, nil

}
func (this DefaultAuthRepository) SignUp(credentials domain.Credentials, user domain.User) (domain.User, error)
