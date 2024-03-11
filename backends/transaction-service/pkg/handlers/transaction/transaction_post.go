package transaction

import (
	"encoding/json"
	"net/http"
	"stori/transaction-service/cmd/config"
	transaction "stori/transaction-service/pkg/application"
)

func HandlerCreateTransaction(dep *config.Dependencies) http.HandlerFunc {

	createTransaction := transaction.NewCreateTransaction(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd transaction.CreateTransactionDto
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		res, err := createTransaction.Exec(r.Context(), &cmd)

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

func HandlerCreateArrayTransaction(dep *config.Dependencies) http.HandlerFunc {

	createTransaction := transaction.NewCreateTransaction(dep.TransactionRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd []transaction.CreateTransactionDto
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		res := make([]transaction.CreateTransactionResponse, len(cmd))
		for _, c := range cmd {
			tran, err := createTransaction.Exec(r.Context(), &c)
			if err != nil {
				config.WriteErr(r.Context(), w, err)
			}
			res = append(res, *tran)
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}
	}

}
