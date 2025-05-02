//Este archivo es el punto de entrada principal de la aplicación. Es el punto de inició de ejecución del programa.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/talialopezarias/WebSite_Go/internal/config"
	"github.com/talialopezarias/WebSite_Go/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	routes.RegisterRoutes()

	addr := fmt.Sprintf(":%s", cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	fmt.Println("Servidor corriendo en http://localhost:" + cfg.Port)

	log.Fatal(server.ListenAndServe())
}
