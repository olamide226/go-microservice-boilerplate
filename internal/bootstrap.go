package bootstrap

import (
	"fmt"

	"github.com/olamide226/go-microservice/pkg/logger"
	"github.com/spf13/viper"
)

func startLogger(environment string) {
	logger.NewLogger(logger.Options{
		Level:       logger.DebugLevel,
		Environment: environment,
	})
}
func SetupDependencies() {
	// get env files
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	startLogger(viper.GetString("ENVIRONMENT"))

	fmt.Println(viper.Get("PORT"))
	fmt.Println(viper.Get("HTTP_PORT"))

}
