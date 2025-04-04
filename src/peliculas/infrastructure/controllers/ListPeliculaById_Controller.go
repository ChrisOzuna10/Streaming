// src/peliculas/infrastructure/controllers/ListPeliculaByIDController.go
package controllers

import (
	"Streaming/src/peliculas/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ListPeliculaByIDController struct {
	findByIDUseCase *application.FindPeliculaByID
}

func NewListPeliculaByIDController(useCase *application.FindPeliculaByID) *ListPeliculaByIDController {
	return &ListPeliculaByIDController{findByIDUseCase: useCase}
}

func (c *ListPeliculaByIDController) Handle(ctx *gin.Context) {
	idStr := ctx.Query("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'id' es obligatorio"})
		return
	}

	idParsed, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || idParsed == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El parámetro 'id' debe ser un número válido"})
		return
	}
	id := uint(idParsed)

	pelicula, err := c.findByIDUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Película no encontrada", "detalle": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pelicula)
}
