package http

import (
	"log"
	"net/http"
	"stori/transaction-service/cmd/config"
	"stori/transaction-service/pkg/handlers"
	"stori/transaction-service/pkg/handlers/transaction"
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
		r.Post("/transaction", transaction.HandlerCreateTransaction(deps))
		r.Post("/transaction/list", transaction.HandlerCreateArrayTransaction(deps))
		r.Get("/transaction/{id}", transaction.HandlerFindTransactionById(deps))
		r.Get("/transaction/user/{userId}", transaction.HandlerFindTransactionByUserId(deps))
		r.Get("/transaction/balance", transaction.HandlerFindBalance(deps))
		r.Get("/transaction/balance/{userId}", transaction.HandlerFindBalanceByUserId(deps))
		r.Get("/transaction/movement/{userId}", transaction.HandlerFindMovementByUserId(deps))
	})

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Println("[" + method + "]: '" + route + "' has " + strconv.Itoa(len(middlewares)) + " middlewares")
		return nil
	})

	return r
}
