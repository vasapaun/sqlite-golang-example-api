#!/usr/bin/env bash

BOOK_ID=${1:-1}

curl -X PUT http://localhost:8080/books/$BOOK_ID \
    -H "Content-Type: application/json" \
    -d '{"title":"Dune Messiah", "author":"Frank Herbert", "year":1969}' | jq
