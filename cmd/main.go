package main

import (
	"github.com/olamide226/go-microservice/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/olamide226/go-microservice/internal"
)

func main() {
	logger.NewLogger(logger.Options{
		Level:       logger.DebugLevel,
		Environment: "development",
	})
	logger.Global.Info("Starting the application...")

	bootstrap.SetupServer()
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