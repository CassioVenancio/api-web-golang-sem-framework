package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Função principal
func main() {
	http.HandleFunc("/", index)       // Rota de acesso
	http.ListenAndServe(":8000", nil) // Cria servidor
}

// Função de recebimento de requisição, encaminhamento e resposta.
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

// Servidor em go simples, localhost:8000
// A função pricipal recebe uma requisição, e envia para a função index.
// A função index retorna uma resposta enviando o index.html, marcado com {{define "Index"}} no arquivo, templates/index.html
