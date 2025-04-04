package controllers

import (
	"Streaming/src/users/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterUserController struct {
	UseCase *application.RegisterUserUseCase
}

func NewRegisterUserController(useCase *application.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{UseCase: useCase}
}

func (c *RegisterUserController) Handle(ctx *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	user, err := c.UseCase.Execute(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado correctamente",
		"user":    user,
	})
}
