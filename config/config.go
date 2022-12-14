package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string //development, staging, production

	HTTPPort    string

	AuthorServiceGrpcHost string
	AuthorServiceGrpcPort string

	ArticleServiceGrpcHost string
	ArticleServiceGrpcPort string

}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "uacademy_article"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.AuthorServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_PORT", ":9001"))

	config.ArticleServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_HOST", "localhost"))
	config.ArticleServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_PORT", ":9001"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
