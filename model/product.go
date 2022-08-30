package model

import "project/m/v2/db"

type Product struct {
	ID          int
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

		p.ID = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amounts = amounts

		products = append(products, p) // append adiciona o proximo elemento para o array, da direita para a esquerda
	}

	defer db.Close() // fecha conecção
	return products
}

func SaveProducts(name string, description string, price float64, amounts int) {
	db := db.ConnectDb()

	query, err := db.Prepare("insert into produtos(name, description, price, amounts) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, description, price, amounts)
	defer db.Close()
}

func DeleteProductById(id string) {
	db := db.ConnectDb()

	deleteQuery, errDelete := db.Prepare("delete from produtos where id=$1")
	if errDelete != nil {
		panic(errDelete.Error())
	}

	deleteQuery.Exec(id)
	defer db.Close()
}
