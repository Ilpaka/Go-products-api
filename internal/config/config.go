package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecret string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: .env not loaded: %v", err)
	}
	return &Config{
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
