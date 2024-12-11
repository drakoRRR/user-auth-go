package main

import (
	"log"
	"os"

	"github.com/drakoRRR/user-auth-go/pkg/config"
	"github.com/drakoRRR/user-auth-go/pkg/migrations"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Envs

	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		if err := migrations.ApplyMigrations(cfg.Database.DSN, config.MigrationsPath, "up"); err != nil {
			log.Fatalf("Migration up failed: %v", err)
		}
		log.Println("Migration up applied successfully.")
	case "down":
		if err := migrations.ApplyMigrations(cfg.Database.DSN, config.MigrationsPath, "down"); err != nil {
			log.Fatalf("Migration down failed: %v", err)
		}
		log.Println("Migration down applied successfully.")
	default:
		log.Printf("Unknown command: %s", cmd)
	}
}
