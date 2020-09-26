package util

import (
	"log"
	"os"
	"strconv"
)

func GetEnvInt(key string, fallback int) int {
	if valueString, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(valueString)
		if err != nil {
			log.Fatalf("Environment variable %s could not be parsed as an integer. Found value: %s. Will use default value: %d\n", key, valueString, fallback)
			return fallback
		}
		return value
	}
	return fallback
}
