package routes

import (
	"github.com/gin-gonic/gin"
)

const (
	// apiv1 prefix for all v1 routes
	ApiV1 string = "/api/v1"
)

func Setup(server *gin.Engine) {
	publicRouter := server.Group(ApiV1)
	// All Public APIs
	HealthCheckRouter(publicRouter)

	// protectedRouter := server.Group(ApiV1)
	// Middleware to verify AccessToken
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	// NewTaskRouter(env, timeout, db, protectedRouter)

	// routes not found
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "endpoint not found"})
	})
}
