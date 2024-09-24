package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	m "API/models"
	"strconv"
)

var Library []m.Book

func InitDataBase() {

	book1 := m.Book{Id: 1, Title: "The Lord of the Rings", Author: "J.R.R. Tolkien"}
	Library = append(Library, book1)
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func Add(c *gin.Context) {

	aStr := c.Param("a")
	bStr := c.Param("b")
	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": a + b})
}

func Multiply(c *gin.Context) {
	aStr := c.Param("a")
	bStr := c.Param("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	resultat := a * b
	c.JSON(http.StatusOK, gin.H{"result": resultat})
}

func Minus(c *gin.Context) {
	aStr := c.Param("a")
	bStr := c.Param("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}
	resultat := a - b
	c.JSON(http.StatusOK, gin.H{"result": resultat})
}
