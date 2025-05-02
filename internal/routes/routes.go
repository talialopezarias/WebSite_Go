// Este archivo registra las rutas y las asocia con sus respectivos manejadores.
package routes

import (
	"net/http" //para manejar rutas

	"github.com/talialopezarias/WebSite_Go/internal/handlers"
)

// Función que registra las rutas para la pagina principal y de error.
func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)

	fs := http.FileServer(http.Dir("web/statics"))

	http.Handle("/statics/", http.StripPrefix("/statics/", fs))
}
