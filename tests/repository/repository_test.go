package repository

import (
	"database/sql"
	"github.com/drakoRRR/user-auth-go/pkg/config"
	"github.com/drakoRRR/user-auth-go/pkg/migrations"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const MigrationsPathTest = "file://../../migrations"

var testDB *sql.DB

func TestMain(m *testing.M) {
	err := migrations.ApplyMigrations(config.Envs.TestDatabase.DSN, MigrationsPathTest, "up")
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	testDB, err = sql.Open("postgres", config.Envs.TestDatabase.DSN)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}
	defer testDB.Close()

	code := m.Run()

	if err := migrations.ApplyMigrations(config.Envs.TestDatabase.DSN, MigrationsPathTest, "down"); err != nil {
		log.Printf("Failed to rollback migrations: %v", err)
	}

	os.Exit(code)
}
