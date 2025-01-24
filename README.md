Using Gin
Get: curl http://localhost:8080/books
Post: curl http://localhost:8080/books --include --header "Content-Type: application/json" -d "@demoData.json" --request POST