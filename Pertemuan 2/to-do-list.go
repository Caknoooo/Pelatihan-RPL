package main

import(
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var db *sql.DB

func main(){
	var err error
	db, err := sql.Open("pgx", "postgres://postgres:Cakno6969@localhost:5432/todo")
	if err != nil{
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/", getUser)
	http.HandleFunc("/users/create", createUser)
	http.HandleFunc("/users/update/", updateUser)
	http.HandleFunc("/users/delete/", deleteUser)

	fmt.Println("Starting server on port http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, r *http.Request){
	rows, err := db.Query("SELECT * FROM users");
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user User
	row := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id)

	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil{
		if err == sql.ErrNoRows{
			http.NotFound(w, r)
		} else{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request){
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	row := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Age)
	if err := row.Scan(&newUser.ID); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[len("/users/update/"):])
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updateUser User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	updateUser.ID = id
	_, err = db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", updateUser.Name, updateUser.Age, updateUser.ID)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated",
	})
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Path[len("/users/delete/"):])
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted",
	})
}