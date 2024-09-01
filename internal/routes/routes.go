package routes

import (
	"go-adAstra/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", controllers.HealthCheck)

	return router
}