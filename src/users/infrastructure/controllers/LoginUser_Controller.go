// src/users/controllers/login_user_controller.go
package controllers

import (
	"Streaming/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginUserController struct {
	UseCase *application.LoginUserUseCase
}

func NewLoginUserController(useCase *application.LoginUserUseCase) *LoginUserController {
	return &LoginUserController{UseCase: useCase}
}

func (c *LoginUserController) Handle(ctx *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	token, err := c.UseCase.Execute(credentials.Email, credentials.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Correo o contraseña inválidos"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
