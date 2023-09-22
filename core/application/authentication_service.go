package application

import (
	"auth_api/core"
	"auth_api/core/domain"
)

type AuthenticationService struct {
	authenticationService domain.AuthenticationService
}

func (this *AuthenticationService) New(AuthenticationService domain.AuthenticationService) {
	this.authenticationService = AuthenticationService
}

func (this *AuthenticationService) SignIn(Credentials core.CredentialsDto) (core.SessionDto, error) {
	session, error := this.authenticationService.Authenticate(domain.Credentials{
		Username: Credentials.Username,
		Password: Credentials.Password,
	})
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
func (this *AuthenticationService) SignOut(Session core.SessionDto) error {
	error := this.authenticationService.Logout(
		domain.Session{
			Id: Session.Id,
		},
	)

	if error != nil {
		return error
	}
	return nil
}
