package main

import (
	"net/http"
	"project/m/v2/routes"
)

// Função principal
func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil) // Cria servidor
}

// Servidor em go simples, localhost:8000
