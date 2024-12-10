package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PublicHost string
	Port       string
	Database   *Database
}

type Database struct {
	Port     string
	User     string
	Password string
	Host     string
	Name     string
	DSN      string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := &Database{
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Host:     getEnv("DB_HOST", "postgres"),
		Name:     getEnv("DB_NAME", "postgres"),
	}

	dbConfig.DSN = fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
	)

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8080"),
		Database:   dbConfig,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
