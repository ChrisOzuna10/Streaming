package routes

import (
	"Streaming/src/users/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, loginCtrl *controllers.LoginUserController, registerCtrl *controllers.RegisterUserController) {
	userGroup := router.Group("/users")

	userGroup.POST("/login", loginCtrl.Handle)
	userGroup.POST("/register", registerCtrl.Handle)
}
