package routes

import (
	"net/http"

	"github.com/drakoRRR/user-auth-go/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func GetRoutes(Handler *handlers.Handlers) []Route {
	return []Route{
		{
			Method:      http.MethodGet,
			Path:        "/health",
			HandlerFunc: Handler.HealthCheck,
		},
		{
			Method:      http.MethodPost,
			Path:        "/users",
			HandlerFunc: Handler.CreateUserHandler,
		},
		{
			Method:      http.MethodGet,
			Path:        "/swagger/*",
			HandlerFunc: httpSwagger.WrapHandler,
		},
	}
}

func RegisterRoutes(mux *http.ServeMux, routes []Route) {
	for _, route := range routes {
		mux.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != route.Method {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			route.HandlerFunc(w, r)
		})
	}
}
