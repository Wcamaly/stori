package transaction

import (
	"encoding/json"
	"net/http"
	"stori/transaction-service/cmd/config"
	transaction "stori/transaction-service/pkg/application"
	"stori/transaction-service/pkg/domain/models"

	"github.com/go-chi/chi/v5"
)

func HandlerFindTransactionById(dep *config.Dependencies) http.HandlerFunc {
	getTransactionById := transaction.NewGetTransactionById(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dto := &transaction.GetTransactionByIdDto{
			ID: models.ID(id),
		}

		res, err := getTransactionById.Exec(r.Context(), dto)
		if err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}
}

func HandlerFindBalanceByUserId(dep *config.Dependencies) http.HandlerFunc {
	getBalanceByUserId := transaction.NewGetBalanceByUserId(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		dto := &transaction.GetBalanceByUserIdDto{
			UserId: models.ID(id),
		}

		res, err := getBalanceByUserId.Exec(r.Context(), dto)
		if err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}
}

func HandlerFindTransactionByUserId(dep *config.Dependencies) http.HandlerFunc {
	getTransactionByUserId := transaction.NewGetTransactionByUserId(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		dto := &transaction.GetTransactionByUserIdDto{
			ID: models.ID(id),
		}

		res, err := getTransactionByUserId.Exec(r.Context(), dto)
		if err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}
}

func HandlerFindMovementByUserId(dep *config.Dependencies) http.HandlerFunc {
	getMovementByUserId := transaction.NewGetMovementByUserId(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")

		dto := &transaction.GetMovementByUserIdDto{
			UserId: models.ID(id),
		}

		res, err := getMovementByUserId.Exec(r.Context(), dto)
		if err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}
}

func HandlerFindBalance(dep *config.Dependencies) http.HandlerFunc {
	getBalance := transaction.NewGetBalance(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := getBalance.Exec(r.Context())
		println("res: ", res)
		if err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}
}
