package controllers

import (
	"Streaming/src/peliculas/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListPeliculasController struct {
	getAllPeliculasUseCase *application.GetAllPeliculas
}

func NewListPeliculasController(useCase *application.GetAllPeliculas) *ListPeliculasController {
	return &ListPeliculasController{getAllPeliculasUseCase: useCase}
}

func (c *ListPeliculasController) Handle(ctx *gin.Context) {
	peliculas, err := c.getAllPeliculasUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las películas: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, peliculas)
}
