package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Init books var as a slice of Books struct
var books []Book

func main() {
	router := mux.NewRouter()

	// Mock Data Implement DB
	books = append(books, Book{ID: "1", Isbn: "42323", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "42332", Title: "Book Two", Author: &Author{FirstName: "Steve", LastName: "Smith"}})
	books = append(books, Book{ID: "3", Isbn: "423425", Title: "Book Three", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "4", Isbn: "42323653", Title: "Book Four", Author: &Author{FirstName: "John", LastName: "Doe"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Listening to port 8000: ")
	log.Fatal(http.ListenAndServe(":8000", router))
}
