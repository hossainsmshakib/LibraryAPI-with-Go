package main

import (
	"github.com/gin-gonic/gin"
)

type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{

	{Id: "1", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 5},
	{Id: "2", Title: "To Kill a Mockingbird", Author: "Harper Lee", Quantity: 3},
	{Id: "3", Title: "1984", Author: "George Orwell", Quantity: 2},
	{Id: "4", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 4},
	{Id: "5", Title: "The Catcher in the Rye", Author: "J.D. Salinger", Quantity: 5},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8080")
}
