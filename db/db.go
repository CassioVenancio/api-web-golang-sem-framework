package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// função para conecção com o banco de dados
func ConnectDb() *sql.DB {
	conn := "user=postgres dbname=store-db password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
