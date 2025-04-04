package routes

import (
	dependencies "Streaming/src/peliculas/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupPeliculasRoutes(router *gin.Engine, deps *dependencies.PeliculaDependencies) {
	peliculaGroup := router.Group("/peliculas")

	{
		peliculaGroup.POST("", deps.CreatePeliculaController.Handle)
		peliculaGroup.PUT("", deps.UpdatePeliculaController.Handle)
		peliculaGroup.DELETE("", deps.DeletePeliculaController.Handle)
		peliculaGroup.GET("", deps.ListPeliculasController.Handle)
		peliculaGroup.GET("/detalle", deps.ListPeliculaByIDController.Handle)
		peliculaGroup.GET("/genero", deps.ListPeliculasByGenreController.Handle)
		peliculaGroup.GET("/buscar", deps.ListPeliculasByGenreController.Handle)
	}
}
