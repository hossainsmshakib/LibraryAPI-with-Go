package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func bookById(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for _, book := range books {
		if book.Id == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}
