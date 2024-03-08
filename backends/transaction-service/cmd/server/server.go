package server

import (
	"fmt"
	"log"
	"net/http"
	"stori/transaction/cmd/config"
	routeHttp "stori/transaction/pkg/ports/http"
)

func StartServer(cfg *config.Config, deps *config.Dependencies) error {
	router := routeHttp.NewRouter(deps)
	log.Println("Servidor escuchando en el puerto %s", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), router)
}