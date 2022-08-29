package controllers

import (
	"net/http"
	"project/m/v2/model"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Função de recebimento de requisição, encaminhamento e resposta.
func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
