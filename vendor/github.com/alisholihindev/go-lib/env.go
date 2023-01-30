package lib

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConfig(key string) string {
	//load file env
	_ = godotenv.Load(".env")

	return os.Getenv(key)
}
