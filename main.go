package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getById", GetBookById).Methods("GET")
	router.HandleFunc("/getBooks", GetBooks).Methods("GET")
	router.HandleFunc("/searchByTitle", SearchByTitle).Methods("GET")
	router.HandleFunc("/getBooksByAsc", SortOrderByAsc).Methods("GET")
	router.HandleFunc("/getBooksByDesc", SortOrderByDesc).Methods("GET")
	router.HandleFunc("/updateById", UpdateById).Methods("PUT")
	router.HandleFunc("/deleteById", DeleteById).Methods("DELETE")
	router.HandleFunc("/addBook", AddBook).Methods("POST")
	
    errr := http.ListenAndServe(":8080", router)
	if errr != nil {
        fmt.Println(errr)
    }
}