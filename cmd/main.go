package main

import (
	"github.com/drakoRRR/user-auth-go/cmd/api"
	"github.com/drakoRRR/user-auth-go/internal/db"
	"github.com/drakoRRR/user-auth-go/internal/handlers"
	"github.com/drakoRRR/user-auth-go/internal/repository"
	"github.com/drakoRRR/user-auth-go/internal/service"
	"github.com/drakoRRR/user-auth-go/pkg/config"
	"github.com/drakoRRR/user-auth-go/pkg/logger"
	"go.uber.org/zap"
	"log"

	_ "github.com/drakoRRR/user-auth-go/docs"
)

func main() {
	cfg := config.Envs

	projectLogger := logger.New()
	projectLogger.Info("Logger initialized successfully")

	dbConn, err := db.ConnectDB(*cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbConn.Close()

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(*userRepo)

	services := &handlers.Services{
		Users: *userService,
	}

	server := api.InitServer(&config.Envs, services, projectLogger)
	projectLogger.Info("Server is running on %s:%s", zap.String("host", cfg.PublicHost), zap.String("port", cfg.Port))
	if err := server.ListenAndServe(); err != nil {
		projectLogger.Fatal("Server failed to start", zap.Error(err))
	}
}