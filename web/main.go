package main

import (
	//Permite imprimir mensajes en la consola
	"fmt"
	//Proporciona funcionalidades para construir servidores web con el protocolo HTTP
	"net/http"
	//Vamos a registrar todo lo que esta ocurriendo dentro del servidor
	"log"
)

func main() {
	// http.HandleFunc espera dos cosas:
	// 1. Una ruta ("/") que será la URL que se escuche
	// 2. Una función que se ejecuta cuando alguien accede a esa ruta.
	//El primero se utiliza para enviar la respuesta al cliente y el Request representa la solicitud entrante.
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Servidor está corriendo en http://localhost:8080")
	//Primer argumento dirección IP y puerto, y el segundo para utilizar el manejador ya predeterminado.
	// "nil" indica que usamos el manejador por defecto que ya configuramos con HandleFunc
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// El primero se utiliza para enviar la respuesta al cliente y el Request representa la solicitud entrante.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Respondemos al navegador escribiendo directamente en el objeto "w"
	// Esto se traduce en que el navegador verá "Hola mundo" en pantalla
	fmt.Fprintln(w, "¡Bienvenido a la página de inicio con Go!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Esta es la página 'Acerca de nosotros'")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Póngase en contácto con nosotros en talialopezarias@gmail.com ")
}
