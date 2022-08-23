package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

// função para conecção com o banco de dados
func connectDb() *sql.DB {
	conn := "user=postgres dbname=store-db password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

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
	db := connectDb()
	stmt, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for stmt.Next() {
		var id, amounts int
		var name, description string
		var price float64

		err = stmt.Scan(&id, &name, &description, &price, &amounts)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Amounts = amounts

		products = append(products, p) // append adiciona o proximo elemento para o array, da direita para a esquerda
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close() // fecha conecção
}

// Servidor em go simples, localhost:8000
