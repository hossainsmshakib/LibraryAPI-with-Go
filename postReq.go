package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}