package util

import (
	"log"
	"os"
)

func GetEnvOrPanic(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Panicf("`%s` environment variable not set.", key)
	}

	return value
}
