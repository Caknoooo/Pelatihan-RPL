package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

type User struct{
	ID uint `json:"id"`
	Name string `json:"user"`
}

func Connect()(*sql.DB, error){
	db, err := sql.Open("pgx", "postgres://postgres:Cakno6969@localhost:5432/pgx")
	if err != nil{
		return nil, err
	}
	return db, err
}

func InsertToDB(db *sql.DB, user User)(*User, error){
	rows, err := db.Query("INSERT INTO users (name) VALUES ($1) RETURNING id, name", user.Name)
	if err != nil{
		return nil, err
	}
	
	rows.Next()

	result := User{}

	rows.Scan(&result.ID, &result.Name)
	return &result, nil
}

func GetAll(db *sql.DB)([]User, error){
	var result []User
	rows, err := db.Query("SELECT * FROM users")

	if err != nil{
		return nil, err
	}

	for rows.Next(){
		var user User
		rows.Scan(&user.ID, &user.Name)
		result = append(result, user)
	}

	return result, nil
}

func main(){
	db, err := Connect()
	if err!=nil{
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil{
		fmt.Println(err)
	}

	res, err := GetAll(db)
	if err!=nil{
		fmt.Println(err)
	}

	jsonMap := map[string]interface{}{
		"data": res,
	}

	b, err := json.Marshal(jsonMap)
	if err != nil{
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Println(w, string(b))
	})

	fmt.Println("Starting Web Server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}