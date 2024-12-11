package handlers

import (
	"net/http"

	"github.com/drakoRRR/user-auth-go/internal/service"
	"github.com/drakoRRR/user-auth-go/pkg/config"
	"github.com/drakoRRR/user-auth-go/pkg/logger"
)

type Services struct {
	Users service.UserService
}

type Handlers struct {
	App      *config.Config
	Services *Services
	Log      *logger.Logger
}

func NewHandlers(app *config.Config, services *Services, log *logger.Logger) *Handlers {
	return &Handlers{
		App:      app,
		Services: services,
		Log:      log,
	}
}

func (h *Handlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.Log.Info("Health check endpoint hit")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
