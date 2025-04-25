package main

import (
	//Permite imprimir mensajes en la consola
	"fmt"
	//Proporciona funcionalidades para construir servidores web con el protocolo HTTP
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Respondemos al navegador escribiendo directamente en el objeto "w"
	// Esto se traduce en que el navegador verá "Hola mundo" en pantalla
	fmt.Fprintln(w, "Hola mundo con Go")

}

func main() {
	// Registramos una función que manejará las solicitudes HTTP a la ruta "/"
	// http.HandleFunc espera dos cosas:
	// 1. Una ruta ("/") que será la URL que se escuche
	// 2. Una función que se ejecuta cuando alguien accede a esa ruta.
	//El primero se utiliza para enviar la respuesta al cliente y el Request representa la solicitud entrante.
	http.HandleFunc("/", handler)
	fmt.Println("Servidor está corriendo en http://localhost:8080")
	//Primer argumento dirección IP y puerto, y el segundo para utilizar el manejador ya predeterminado.
	// "nil" indica que usamos el manejador por defecto que ya configuramos con HandleFunc
	http.ListenAndServe(":8080", nil)
}
