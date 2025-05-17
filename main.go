package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllBooks(w, r)
		case http.MethodPost:
			addBook(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBook(w, r)
		case http.MethodPut:
			updateBook(w, r)
		case http.MethodDelete:
			deleteBook(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("API running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
