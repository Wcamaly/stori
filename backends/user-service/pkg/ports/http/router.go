package http

import (
	"log"
	"net/http"
	"stori/user-service/cmd/config"
	"stori/user-service/pkg/handlers"
	"stori/user-service/pkg/handlers/user"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter configura y retorna un nuevo enrutador
func NewRouter(deps *config.Dependencies) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Configurar rutas
	r.Get("/health", handlers.StatusHealth(deps))
	r.Get("/openapi-json", handlers.SwaggerHandler(deps))

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/user", user.HandlerCreateUser(deps))
		r.Get("/user/{id}", user.HandlerFindUserById(deps))
		r.Get("/user/email/{email}", user.HandlerFindUserByEmail(deps))
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println("[" + method + "]: '" + route + "' has " + strconv.Itoa(len(middlewares)) + " middlewares")
		return nil
	})

	return r
}
