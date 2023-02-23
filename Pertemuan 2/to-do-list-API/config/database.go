package config

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase()(*sql.DB, error){
	db, err := sql.Open("pgx", "postgresql://postgres:root@localhost:5432/todolist_rpl")
	if err != nil{
		return nil, err
	}

	fmt.Println("Database Connected")
	return db, nil
}

