package handlers

import (
	"io/ioutil"
	"net/http"
	"stori/user-service/cmd/config"
)

func SwaggerHandler(deps *config.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		swaggerJSON, err := ioutil.ReadFile("swagger.json")
		if err != nil {
			http.Error(w, "error reading swagger.json file", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(swaggerJSON)
	}

}
