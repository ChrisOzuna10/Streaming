package controllers

import (
	"Streaming/src/peliculas/application"
	"Streaming/src/peliculas/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdatePeliculaController struct {
	updatePeliculaUseCase *application.UpdatePelicula
}

func NewUpdatePeliculaController(useCase *application.UpdatePelicula) *UpdatePeliculaController {
	return &UpdatePeliculaController{updatePeliculaUseCase: useCase}
}

func (c *UpdatePeliculaController) Handle(ctx *gin.Context) {
	var pelicula entities.Pelicula
	if err := ctx.ShouldBindJSON(&pelicula); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar datos: " + err.Error()})
		return
	}

	if pelicula.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de la película es requerido"})
		return
	}

	updated, err := c.updatePeliculaUseCase.Execute(pelicula)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la película: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}
