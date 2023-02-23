package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("pgx", "postgresql://postgres:Cakno6969@localhost:5432/pgx")
	if err != nil {
		return nil, err
	}

	return db, err
}

func main() {
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}