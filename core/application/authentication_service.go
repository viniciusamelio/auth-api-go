package application

import (
	"auth_api/core"
	"auth_api/core/domain"
)

type ApplicationAuthenticationService interface {
	SignIn(Credentials core.CredentialsDto) (core.SessionDto, error)
	SignOut(Session core.SessionDto) (bool, error)
	SignUp(Data core.SignUpDto) (core.UserDto, error)
}

type AuthenticationService struct {
	authenticationService domain.DomainAuthenticationService
}

func (this *AuthenticationService) New(AuthenticationService domain.DomainAuthenticationService) {
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
func (this *AuthenticationService) SignOut(Session core.SessionDto) (bool, error) {
	_, error := this.authenticationService.Logout(
		domain.Session{
			Id: Session.Id,
		},
	)

	if error != nil {
		return false, error
	}
	return true, nil
}

func (this *AuthenticationService) SignUp(Data core.SignUpDto) (core.UserDto, error) {
	user, error := this.authenticationService.SignUp(domain.Credentials{
		Username: Data.Email,
		Password: Data.Password,
	}, domain.User{
		Name:  Data.Name,
		Email: Data.Email,
	})

	if error != nil {
		return core.UserDto{}, error
	}

	return core.UserDto{
		Email: user.Email,
		Name:  user.Name,
		Id:    user.Id,
	}, nil
}
