# SQLite3 Golang Book API

This is a simple, example API to manage books (title, author, year) with SQLite3 and Golang

Tests are written in bash, using curl

# How to use

To run:

`go run main.go db.go handlers.go model.go` 

To compile and run:

`go build && ./book-api`

# How to test

cd into tests/ and run any of the bash scripts (some use jq to parse json)

For delete_book, update_book, and get_book, provide a book id as the only argument.
