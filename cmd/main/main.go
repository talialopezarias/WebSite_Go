//Este archivo es el punto de entrada principal de la aplicación. Es el punto de inició de ejecución del programa.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/talialopezarias/WebSite_Go/internal/config"
	"github.com/talialopezarias/WebSite_Go/internal/routes"
)

func main() {
	logFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error al abrir el archivo de registro: ", err)
		return
	}

	defer logFile.Close() //aseguramos que el archivo se cierre correctamente cuando main se cierra

	log.SetOutput(logFile)

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
