// src/users/infrastructure/dependencies.go
package dependencies

import (
	"Streaming/src/users/application"
	"Streaming/src/users/infrastructure/controllers"
	infrastructure "Streaming/src/users/infrastructure/databases"
)

type UserDependencies struct {
	LoginController    *controllers.LoginUserController
	RegisterController *controllers.RegisterUserController
}

func InitUserDependencies() *UserDependencies {
	repo := infrastructure.NewMySQLUserRepository()

	loginUC := application.NewLoginUserUseCase(repo)
	registerUC := application.NewRegisterUserUseCase(repo)

	return &UserDependencies{
		LoginController:    controllers.NewLoginUserController(loginUC),
		RegisterController: controllers.NewRegisterUserController(registerUC),
	}
}
