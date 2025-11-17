package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                   string
	DBURI                  string
	HostName               string
	JWTSecret              string
	APIVersion             string
	JWTExpirationInSeconds int64
}

var Envs = initializeConfig()

func initializeConfig() Config {
	godotenv.Load()

	return Config{
		Port:                   getEnv("PORT", ""),
		DBURI:                  getEnv("DB_URI", ""),
		HostName:               getEnv("HOSTNAME", ""),
		JWTSecret:              getEnv("JWT_SECRET", ""),
		APIVersion:             getEnv("API_VERSION", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
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
