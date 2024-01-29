package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olamide226/go-microservice/internal/bootstrap"
	"github.com/olamide226/go-microservice/internal/routes"
	"github.com/olamide226/go-microservice/pkg/logger"
)

func main() {

	env := bootstrap.SetupDependencies()

	logger.Global.Info("Starting the application...")

	router := gin.Default()
	routes.Setup(router)
	router.Run(env.ServerAddress) // listen and serve on
}
