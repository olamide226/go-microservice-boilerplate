package bootstrap

import (
	"github.com/olamide226/go-microservice/pkg/logger"
)

func startLogger(environment string) {
	logger.NewLogger(logger.Options{
		Environment: environment,
	})
}
func SetupDependencies() *Env {
	env := NewEnv()

	startLogger(env.Environment)
	return env
}
