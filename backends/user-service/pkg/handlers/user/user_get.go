package user

import (
	"encoding/json"
	"net/http"
	"stori/user-service/cmd/config"
	"stori/user-service/pkg/application/user"

	"github.com/go-chi/chi/v5"
)

func HandlerFindUserByEmail(dep *config.Dependencies) http.HandlerFunc {

	getUserByEmail := user.NewGetUserByEmail(dep.UserRepository)

	return func(writer http.ResponseWriter, request *http.Request) {

		email := chi.URLParam(request, "email")

		res, err := getUserByEmail.Exec(request.Context(), &user.GetUserByEmailRequest{Email: email})

		if err != nil {
			config.WriteErr(request.Context(), writer, err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(res); err != nil {
			config.WriteErr(request.Context(), writer, err)
			return
		}
	}
}

func HandlerFindUserById(dep *config.Dependencies) http.HandlerFunc {
	getUserById := user.NewGetUserById(dep.UserRepository)
	return func(writer http.ResponseWriter, request *http.Request) {
		id := chi.URLParam(request, "id")
		res, err := getUserById.Exec(request.Context(), &user.GetUserRequestById{ID: id})

		if err != nil {
			config.WriteErr(request.Context(), writer, err)
			return
		}

		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(res); err != nil {
			config.WriteErr(request.Context(), writer, err)
			return
		}

	}
}
