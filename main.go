package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amounts     int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Função principal
func main() {
	http.HandleFunc("/", index)       // Rota de acesso
	http.ListenAndServe(":8000", nil) // Cria servidor
}

// Função de recebimento de requisição, encaminhamento e resposta.
func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Shirt", Description: "No description", Price: 3, Amounts: 40},
		{Name: "Pants", Description: "No description", Price: 5, Amounts: 6},
	}

	temp.ExecuteTemplate(w, "Index", products)
}

// Servidor em go simples, localhost:8000
