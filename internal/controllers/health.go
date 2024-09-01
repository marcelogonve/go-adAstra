package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 	http.StatusOK,
		"message": 	"AdAstra API is running 🚀",
	})
}