#!/usr/bin/env bash

BOOK_ID=${1:-1}
curl -s http://localhost:8080/books/$BOOK_ID | jq
