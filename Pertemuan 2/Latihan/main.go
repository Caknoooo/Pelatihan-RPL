package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Book struct{
	ISDN string `json: "isdn"`
	Title string `json: "title"`
	Author string `json: "author"`
	Pages int `json: "pages"`
}

type JsonResponse struct{
	Meta interface{} `json: "meta"`
	Data interface{} `json: "data"`
}

type ApiError struct{
	Status int16 `json: "status"`
	Title string `json: "title`
}

type JsonResponseError struct{
	Error *ApiError `json: "error"`
}

var bookStore = make(map[string] *Book)

func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	books := []*Book{}

	for _, book := range bookStore{
		books = append(books, book)
	}

	response := &JsonResponse{Data: &books}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil{
		panic(err)
	}
}

func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	isdn := params.ByName("isdn")
	book, ok := bookStore[isdn]
	w.Header().Set("Content-Type", "application/json");

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		response := JsonResponseError{Error: &ApiError{Status: 404, Title: "Web"}}

		if err := json.NewEncoder(w).Encode(response); err != nil{
			panic(err)
		}
	}
	response := JsonResponse{Data: book}
	if err := json.NewEncoder(w).Encode(response); err != nil{
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcome");
}

func main(){
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/books", BookIndex)
	router.GET("/books/:isdn", BookShow)

	bookStore["123"] = &Book{
		ISDN: "123",
		Title: "Thomas and tompel",
		Author: "Thomas hehe",
		Pages: 367,
	}

	bookStore["124"] = &Book{
		ISDN: "124",
		Title: "Thomas and Caknoo",
		Author: "Caknoo hehe",
		Pages: 367,
	}

	fmt.Print("Port running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}