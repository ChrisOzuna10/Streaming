package application

import (
	"Streaming/src/users/domain"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginUserUseCase struct {
	Repo domain.IUser
}

func NewLoginUserUseCase(repo domain.IUser) *LoginUserUseCase {
	return &LoginUserUseCase{Repo: repo}
}

func (uc *LoginUserUseCase) Execute(email, password string) (string, error) {
	user, err := uc.Repo.FindByEmail(email)
	if err != nil || user.ID == 0 {
		return "", errors.New("correo o contraseña inválidos")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("correo o contraseña inválidos")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("clave_secreta_super_segura")

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
