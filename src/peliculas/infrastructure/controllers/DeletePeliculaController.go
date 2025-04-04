package controllers

import (
	"Streaming/src/peliculas/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DeletePeliculaController struct {
	deletePeliculaUseCase *application.DeletePelicula
}

func NewDeletePeliculaController(useCase *application.DeletePelicula) *DeletePeliculaController {
	return &DeletePeliculaController{deletePeliculaUseCase: useCase}
}

func (c *DeletePeliculaController) Handle(ctx *gin.Context) {
	idStr := ctx.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.deletePeliculaUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la película: " + err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}
