//Este archivo registra las rutas y las asocia con sus respectivos manejadores.

package routes

import (
	"net/http" //para manejar rutas
)

// Función que registra las rutas
func RegisterRoutes() {
	http.HandleFunc("/", handlers.homeHandler)
	http.HandleFunc("/error", handlers.errorHandler)

	fs := http.FileServer(http.Dir("web/statics"))

	http.Handle("/statics/", http.StripPrefix("/statics/", fs))
}
