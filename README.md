Using Gin
Get: http://localhost:8080/books
Post: http://localhost:8080/books (to add a new book)
    - curl http://localhost:8080/books --include --header "Content-Type: application/json" -d "@demoData.json" --request POST