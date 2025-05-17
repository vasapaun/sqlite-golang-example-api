package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	row := DB.QueryRow("SELECT * FROM books WHERE id = ?", id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	res, _ := DB.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
	id, _ := res.LastInsertId()
	book.ID = int(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")

	intid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if !bookExists(intid) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	_, err = DB.Exec("UPDATE books SET title=?, author=?, year=? WHERE id=?", book.Title, book.Author, book.Year, id)
	if err != nil {
		http.Error(w, "Update failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")

	intid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if !bookExists(intid) {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	row := DB.QueryRow("SELECT * FROM books WHERE id = ?", id)
	if row == nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	_, err = DB.Exec("DELETE from books WHERE id=?", id)
	if err != nil {
		http.Error(w, "Deletion failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func bookExists(id int) bool {
	var exists bool

	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
