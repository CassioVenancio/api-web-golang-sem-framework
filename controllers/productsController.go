package controllers

import (
	"log"
	"net/http"
	"project/m/v2/model"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Função de recebimento de requisição, encaminhamento e resposta.
func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := model.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		desc := r.FormValue("description")
		price, errPriceConv := strconv.ParseFloat(r.FormValue("price"), 64)
		amounts, errAmountsConv := strconv.Atoi(r.FormValue("amounts"))

		if errPriceConv != nil {
			log.Println("Conversion failed in price", errPriceConv)
		}
		if errAmountsConv != nil {
			log.Println("Conversion failed in amounts", errAmountsConv)
		}

		model.SaveProducts(name, desc, price, amounts)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	model.DeleteProductById(id)
	http.Redirect(w, r, "/", 301)
}
