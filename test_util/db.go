package test_util

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadTestEnv(basePath string) {
	err := godotenv.Load(fmt.Sprintf("%s.env.test", basePath))
	if err != nil {
		panic("No .env.test file found.")
	}
}
