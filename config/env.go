package config

import (
	"os"
	"log"
	dotenv "github.com/joho/godotenv"
)

func EnvLoad() {
	var env_file string = "sample.env"
	if err := dotenv.Load(env_file); err != nil {
		log.Fatalf("Failed to load the file %s: %v", env_file, err)
	}
}

func EnvGet(key string) string {
	return os.Getenv(key)
}