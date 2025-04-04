// src/peliculas/infrastructure/controllers/ListPeliculasByGenreController.go
package controllers

import (
	"Streaming/src/peliculas/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListPeliculasByGenreController struct {
	useCase *application.ListPeliculasByGenre
}

func NewListPeliculasByGenreController(useCase *application.ListPeliculasByGenre) *ListPeliculasByGenreController {
	return &ListPeliculasByGenreController{useCase: useCase}
}

func (c *ListPeliculasByGenreController) Handle(ctx *gin.Context) {
	genre := ctx.Query("genero")
	if genre == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Género no especificado"})
		return
	}

	peliculas, err := c.useCase.Execute(genre)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, peliculas)
}
