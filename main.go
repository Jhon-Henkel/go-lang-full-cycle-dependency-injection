package main

import (
	"database/sql"
	"fmt"
	"github.com/Jhon-Henkel/go-lang-full-cycle-dependency-injection/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := product.NewProductRepository(db)
	useCase := product.NewProductUseCase(repository)

	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
