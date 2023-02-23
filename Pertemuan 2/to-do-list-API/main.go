package main

import (
	"fmt"
	"net/http"

	"github.com/Caknoooo/to-do-list-api/routes"
)

func main(){
	route := routes.Router()	

	fmt.Println("Starting server on the port http://localhost:8080/")
	http.ListenAndServe(":8080", route)
}