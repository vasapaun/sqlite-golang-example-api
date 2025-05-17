#!/usr/bin/env bash

BOOK_ID=${1:-1}
curl -X DELETE http://localhost:8080/books/$BOOK_ID
