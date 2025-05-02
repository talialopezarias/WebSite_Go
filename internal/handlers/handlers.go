//Este archivo contiene los manejadores que gestionan las solicitudes y respuestas del servidor web.
//Todo esto para cargar plantillas HTML y que proporcionan datos dinámicos para su rendederización.

package handlers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/talialopezarias/WebSite_Go/internal/models"
)

func renderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	tmpl, err := template.ParseFiles(tmplFile)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:   "Mi primera página en Go",
		Author:  "Talía López",
		Welcome: "Go es un lenguaje que impactará en tus proyectos webs actuales y futuros.",
	}

	page := r.URL.Path[1:]

	if page == "" {
		page = "index.html"
	}

	tmplFile := "web/templates/" + page

	if _, err := os.Stat(tmplFile); err != nil {
		tmplFile = "web/templates/error.html"

		data.ErrorCode = http.StatusNotFound
		data.ErrorMessage = "Página no encontrada"
	}

	renderTemplate(w, tmplFile, data)

}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title:        "Página no encontrada",
		ErrorCode:    http.StatusInternalServerError,
		ErrorMessage: "Error interno del servidor",
	}

	renderTemplate(w, "web/tempaltes/error.html", data)

}
