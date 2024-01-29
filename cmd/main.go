package main

import (
	"github.com/olamide226/go-microservice/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/olamide226/go-microservice/internal"
)

func main() {
	
	bootstrap.SetupDependencies()

	logger.Global.Info("Starting the application...")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		logger.Global.Debug("Ping endpoint called")
		c.JSON(200, gin.H{
			"message": "pong",
			"level": logger.DebugLevel,
		})
	})
	router.Run(":8000") // listen and serve on
}