package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	useCase := NewUseCase(db)

	prod, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(prod.Name)
}
