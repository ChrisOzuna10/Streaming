// src/users/application/register_user.go
package application

import (
	"Streaming/src/users/domain"
	"Streaming/src/users/domain/entities"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserUseCase struct {
	Repo domain.IUser
}

func NewRegisterUserUseCase(repo domain.IUser) *RegisterUserUseCase {
	return &RegisterUserUseCase{Repo: repo}
}

func (uc *RegisterUserUseCase) Execute(name, email, password string) (entities.User, error) {
	existingUser, _ := uc.Repo.FindByEmail(email)
	if existingUser.ID != 0 {
		return entities.User{}, errors.New("el correo ya está en uso")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}

	user := entities.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	createdUser, err := uc.Repo.Save(user)
	if err != nil {
		return entities.User{}, err
	}

	return createdUser, nil
}
