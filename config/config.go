package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func NewConfig() *Config {

	envFile := ".env"

	injectedEnvFile := os.Getenv("ENV_FILE")
	if injectedEnvFile != "" {
		envFile = injectedEnvFile
	}

	err := godotenv.Load(envFile)
	if err != nil {
		fmt.Println(err)
	}

	config := &Config{
		Port: os.Getenv("PORT"),
	}

	if config.Port == "" {
		config.Port = "8000"
	}

	return config
}
