package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost   string
	Port         string
	Database     *Database
	TestDatabase *Database
}

type Database struct {
	Port     string
	User     string
	Password string
	Host     string
	Name     string
	DSN      string
}

const MigrationsPath = "file://migrations"

var Envs = initConfig()

// LoadENV Added for tests
func LoadENV() {
	path := ".env"
	for {
		err := godotenv.Load(path)
		if err == nil {
			break
		}
		path = "../" + path
	}
}

func initConfig() Config {
	LoadENV()

	dbConfig := &Database{
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Host:     getEnv("DB_HOST", "postgres"),
		Name:     getEnv("DB_NAME", "postgres"),
	}

	dbConfig.DSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	testDbConfig := &Database{
		Port:     getEnv("DB_PORT_TEST", "5433"),
		User:     getEnv("DB_USER_TEST", "postgres"),
		Password: getEnv("DB_PASSWORD_TEST", "postgres"),
		Host:     getEnv("DB_HOST_TEST", "postgres"),
		Name:     getEnv("DB_NAME_TEST", "postgres"),
	}

	testDbConfig.DSN = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		testDbConfig.User,
		testDbConfig.Password,
		testDbConfig.Host,
		testDbConfig.Port,
		testDbConfig.Name,
	)

	return Config{
		PublicHost:   getEnv("PUBLIC_HOST", "http://localhost"),
		Port:         getEnv("PORT", "8080"),
		Database:     dbConfig,
		TestDatabase: testDbConfig,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
