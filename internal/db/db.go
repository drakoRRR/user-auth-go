package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/drakoRRR/user-auth-go/pkg/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Database) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to the database successfully.")
	return db, nil
}
