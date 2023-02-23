package main 

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Ini index");
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hello wolrd!");
	})

	http.HandleFunc("/index", index);

	fmt.Println("Starting web server at http://localhost:8080/");
	http.ListenAndServe(":8080", nil)
}