package application

import (
	"auth_api/config/database"
	"auth_api/core"
	domain "auth_api/core/domain"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type DefaultAuthRepository struct {
	Database database.GoOrmDatabase
}

func (this *DefaultAuthRepository) SignIn(credentials domain.Credentials) (domain.User, error) {
	var userSchema database.User
	result := this.Database.Select("id", "hash").Where("email = ?", credentials.Username).First(&userSchema)
	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("User not found")
		return domain.User{}, defaultError
	}

	error := bcrypt.CompareHashAndPassword([]byte(userSchema.Hash), []byte(credentials.Password))
	if error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("User not found")
		return domain.User{}, defaultError
	}

	this.Database.Select("id", "name", "email").Where("id = ?", userSchema.Id).First(&userSchema)

	return domain.User{
		Id:    userSchema.Id,
		Name:  userSchema.Name,
		Email: userSchema.Email,
	}, nil
}
func (this *DefaultAuthRepository) SignUp(credentials domain.Credentials, user domain.User) (domain.User, error) {
	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if error != nil {
		return domain.User{}, error
	}
	userDto := database.User{
		Email:     credentials.Username,
		Name:      user.Name,
		Hash:      string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Id:        uuid.NewString(),
	}

	result := this.Database.Save(
		&userDto,
	)
	if result.Error != nil {
		defaultError := core.DefaultError{}
		defaultError.SetMessage("User could not be created")
		return domain.User{}, defaultError
	}

	return domain.User{
		Id:    userDto.Id,
		Name:  userDto.Name,
		Email: userDto.Email,
	}, nil
}
