package main

import (
	"Streaming/src/core"

	peliculaDeps "Streaming/src/peliculas/infrastructure"
	peliculaRoutes "Streaming/src/peliculas/infrastructure/routes"

	userDeps "Streaming/src/users/infrastructure"
	userRoutes "Streaming/src/users/infrastructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	// Conexión a la base de datos
	core.ConnectToDataBase()

	// Inicializar dependencias
	users := userDeps.InitUserDependencies()
	peliculas := peliculaDeps.InitPeliculaDependencies()

	// Crear instancia de Gin
	router := gin.Default()

	// Configurar CORS (ajústalo según tu frontend)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rutas
	userRoutes.SetupUserRoutes(router, users.LoginController, users.RegisterController)
	peliculaRoutes.SetupPeliculasRoutes(router, peliculas)

	// Puerto del servidor
	port := ":8080"
	log.Printf("Servidor corriendo en http://localhost%s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
