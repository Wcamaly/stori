package http

import (
	"stori/transaction/cmd/config"
	"stori/transaction/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter configura y retorna un nuevo enrutador
func NewRouter(deps *config.Dependencies) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Configurar rutas
	router.Get("/health", handlers.StatusHealth(deps))

	return router
}
