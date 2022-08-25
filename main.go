package main

import (
	"html/template"
	"net/http"
	"project/m/v2/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Função principal
func main() {
	http.HandleFunc("/", index)       // Rota de acesso
	http.ListenAndServe(":8000", nil) // Cria servidor
}

// Função de recebimento de requisição, encaminhamento e resposta.
func index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

// Servidor em go simples, localhost:8000
