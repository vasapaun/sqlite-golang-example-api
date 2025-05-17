#!/usr/bin/env bash

curl -X POST http://localhost:8080/books \
    -H "Content-Type: application/json" \
    -d '{"title":"Dune", "author":"Frank Herbert", "year":1965}' | jq
