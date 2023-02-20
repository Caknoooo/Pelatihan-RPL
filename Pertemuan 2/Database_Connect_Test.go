package main

import (
    "database/sql"
    "fmt"

    _ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
    ID   uint64
    Name string
}

func Connect() (*sql.DB, error) {
    db, err := sql.Open("pgx", "postgresql://postgres:Cakno6969@localhost:5432/pgx")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func InsertToDB(db *sql.DB, user User) (*User, error) {
    rows, err := db.Query("INSERT INTO users (name) VALUES ($1) RETURNING id, name", user.Name)
    if err != nil {
        return nil, err
    }

    // call rows.Next() to move pointer to first result set
    rows.Next()

    // result container object
    result := User{}

    // insert rows to result
    rows.Scan(&result.ID, &result.Name)
    return &result, nil
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

    caknii := User{
        Name: "caknii",
    }

    res, err := InsertToDB(db, caknii)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(res)
}