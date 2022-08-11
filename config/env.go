package config

import (
	"fmt"
	"os"
)

var (
	PORT           = getEnv("PORT", "3000")
	DB_CONN_STRING = getEnv("DB_CONN_STRING", "")
)

func getEnv(name, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf("Environment variable not found : %v", name))
}
