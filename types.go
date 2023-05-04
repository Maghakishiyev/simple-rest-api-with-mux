package main

// Book Struct Model

type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}
