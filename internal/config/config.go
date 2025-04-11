package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort       string `env:"HTTP_PORT" env-default:"8080"`
	DBHost         string `env:"DB_HOST" env-default:"localhost"`
	DBPort         string `env:"DB_PORT" env-default:"5432"`
	DBUser         string `env:"DB_USER" env-default:"bookuser"`
	DBPassword     string `env:"DB_PASSWORD" env-default:"bookpass"`
	DBName         string `env:"DB_NAME" env-default:"bookdb"`
	JWTSecret      string `env:"JWT_SECRET" env-default:"secret"`
	UserServiceURL string `env:"USER_SERVICE_URL" env-default:"http://localhost:8081"`
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	var cfg Config
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return &cfg
}