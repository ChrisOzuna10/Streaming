package controllers

import (
	"Streaming/src/peliculas/application"
	"Streaming/src/peliculas/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreatePeliculaController struct {
	createPeliculaUseCase *application.CreatePelicula
}

func NewCreatePeliculaController(useCase *application.CreatePelicula) *CreatePeliculaController {
	return &CreatePeliculaController{createPeliculaUseCase: useCase}
}

func (c *CreatePeliculaController) Handle(ctx *gin.Context) {
	var pelicula entities.Pelicula

	if err := ctx.ShouldBindJSON(&pelicula); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar la solicitud: " + err.Error()})
		return
	}

	if pelicula.UserID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El campo user_id es obligatorio y debe ser un número válido"})
		return
	}

	createdPelicula, err := c.createPeliculaUseCase.Execute(pelicula)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la película: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":  "Película creada correctamente",
		"pelicula": createdPelicula,
	})
}
