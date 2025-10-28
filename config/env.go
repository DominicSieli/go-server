package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port				   string
	DBURI				   string
	DBName				   string
	DBUser				   string
	HostName			   string
	DBAddress			   string
	JWTSecret			   string
	APIVersion			   string
	DBPassword			   string
	JWTExpirationInSeconds int64
}

var Envs = initializeConfig()

func initializeConfig() Config {
	godotenv.Load()

	return Config{
		DBURI:					getEnv("DB_URI", ""),
		Port:					getEnv("PORT", ":8080"),
		DBUser:					getEnv("DB_USER", "root"),
		HostName:				getEnv("HOSTNAME", "localhost"),
		APIVersion:				getEnv("API_VERSION", "/api/v1"),
		DBPassword:				getEnv("DB_PASSWORD", "mypassword"),
		DBName:					getEnv("DB_NAME", "go-server-template"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:				getEnv("JWT_SECRET", "not secret anymore?"),
		DBAddress:				fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)

		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
