package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getBooks(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get request to Books endpoint")
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func getBook(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get request to Books/{id} endpoint")

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	//Loop through books and find with id

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
		}
	}
}

func createBook(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("POST request to Books endpoint")

	writer.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe

	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func updateBook(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("PUT request to Books/{id} endpoint")

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]

			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
}

func deleteBook(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Delete request to Books/{id} endpoint")

	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(books)
}
