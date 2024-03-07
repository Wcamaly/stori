package handlers

import (
	"fmt"
	"net/http"
	"stori/transaction/cmd/config"
	response "stori/transaction/pkg/config/http"
)

// Status Health
func StatusHealth(deps *config.Dependencies) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Hola, este es un endpoint de ejemplo!")
		res := [2]int{1, 2}
		response.OK(w, res)
	}
}
