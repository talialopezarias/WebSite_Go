package main

import (
	//Permite imprimir mensajes en la consola
	"fmt"
	//Proporciona funcionalidades para construir servidores web con el protocolo HTTP
	"net/http"
	//Vamos a registrar todo lo que esta ocurriendo dentro del servidor
	"log"
	//Se utiliza para trabajar con rutas de archivos de forma segura en diferentes SO
	"html/template"
	"path/filepath"
)

type PageData struct {
	Title        string
	Message      template.HTML
	ErrorCode    int
	ErrorMessage string
}

func renderTemplate(w http.ResponseWriter, tmplFile string, data PageData) {
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
	data := PageData{
		Title:   "Mi primera página en Go",
		Message: template.HTML("Es una página sencilla pero vamos aprendiendo <b>poco a poco</b> impartida por la programadora Talía López"),
	}

	tmplFile := filepath.Join("web/templates/", "index.html")
	renderTemplate(w, tmplFile, data)

}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:        "Error 404",
		ErrorCode:    404,
		ErrorMessage: "¡Página no encontrada!",
	}

	tmplFile := filepath.Join("web/templates/", "error.html")
	renderTemplate(w, tmplFile, data)

}

func main() {
	//Creamos un manejador de archivos estáticos utilizando fileserver.
	//http.Dir especifica en el directorio raiz donde se encuentran los archivos estaticos
	//Handle es una ruta static para manejar las solicitud de archivos statics y luego utiliza StripPrefix eliminar el prefijo static.
	fs := http.FileServer(http.Dir("web/statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/error", errorHandler)

	//Es una función controladora para la ruta raiz. La función utiliza el ServeFile y lo que hace es enviar el archivo home como respuesta al cliente.
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//http.ServeFile(w, r, filepath.Join("web/templates/", "home.html"))

	//})

	//http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
	//	http.ServeFile(w, r, filepath.Join("web/templates/", "error.html"))

	//})

	fmt.Println("Servidor está corriendo en http://localhost:8080")
	//Primer argumento dirección IP y puerto, y el segundo para utilizar el manejador ya predeterminado.
	// "nil" indica que usamos el manejador por defecto que ya configuramos con HandleFunc
	log.Fatal(http.ListenAndServe(":8080", nil))
}
