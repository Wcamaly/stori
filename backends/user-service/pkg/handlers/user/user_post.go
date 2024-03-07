package user

import (
	"encoding/json"
	"net/http"
	"stori/user-service/cmd/config"
	"stori/user-service/pkg/application/user"
)

func HandlerCreateUser(dep *config.Dependencies) http.HandlerFunc {
	service := user.NewCreateContract(dep.UserRepository)
	return func(w http.ResponseWriter, r *http.Request) {
		var cmd user.CreateUserDto
		if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
			config.WriteErr(r.Context(), w, err)
			return
		}

		res, err := service.Exec(r.Context(), &cmd)

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
