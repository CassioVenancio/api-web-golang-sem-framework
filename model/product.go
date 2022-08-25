package model

import "project/m/v2/db"

type Product struct {
	Name        string
	Description string
	Price       float64
	Amounts     int
}

func FindAllProducts() []Product {
	db := db.ConnectDb()
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

	defer db.Close() // fecha conecção
	return products
}
