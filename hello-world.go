package main

import (
	"net/http"
)

func main() {
	//Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	//Iniciar el servidor
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Juancarlo"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de contacto"))
}
