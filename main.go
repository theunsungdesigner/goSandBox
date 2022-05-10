package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"ls
	"github.com/gorilla/mux"
)

// Init books var as a slice Book Struct
var books []Book

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}


// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create Single Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

//update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	r := mux.NewRouter()

	// TODO create a mockDatabase
	books = append(books, Book{ID: "1", Isbn: "1606", Author: &Author{FirstName: "Cade", LastName: "Thacker"}})
	books = append(books, Book{ID: "2", Isbn: "2301", Author: &Author{FirstName: "Jonathon", LastName: "Wilson"}})
	books = append(books, Book{ID: "3", Isbn: "1748", Author: &Author{FirstName: "Deivid", LastName: "Rodriguez"}})

	// Mock Data
	// var newBook = Book{ID: "4", Isbn: "1234", Title: "johnny", Author: &Author{FirstName: "someDude", LastName: "Dudette"}}
	// var newBook2 = Book{ID: "5", Isbn: "333", Author: &Author{FirstName: "Ray", LastName: "Mond"}}
	

	//Route Handlers /Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("api/books", createBook).Methods("POST")
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("api/books/{id}", deleteBook).Methods("DELETE")
	fmt.Println("we are here")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
