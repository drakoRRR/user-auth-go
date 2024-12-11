package api

import (
	"fmt"
	"net/http"

	"github.com/drakoRRR/user-auth-go/internal/handlers"
	"github.com/drakoRRR/user-auth-go/internal/handlers/routes"
	"github.com/drakoRRR/user-auth-go/pkg/config"
	"github.com/drakoRRR/user-auth-go/pkg/logger"
)

func InitServer(appConfig *config.Config, services *handlers.Services, log *logger.Logger) *http.Server {
	handler := handlers.NewHandlers(appConfig, services, log)

	mux := http.NewServeMux()

	routesList := routes.GetRoutes(handler)
	routes.RegisterRoutes(mux, routesList)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", appConfig.PublicHost, appConfig.Port),
		Handler: mux,
	}

	return server
}
