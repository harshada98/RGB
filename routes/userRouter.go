package routes

import (
	"RGB/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.getUsers())
	incomingRoutes.GET("/users/user_id", controller.getUser())
}
