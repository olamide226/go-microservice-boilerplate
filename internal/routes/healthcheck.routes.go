package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olamide226/go-microservice/pkg/logger"
)

func HealthCheckRouter(router *gin.RouterGroup) {
	router.GET("/healthcheck", func(c *gin.Context) {
		logger.Global.Debug("Healthcheck endpoint called")
		c.JSON(200, gin.H{
			"message": "ok",
			"level":   logger.DebugLevel,
		})
	})
}
