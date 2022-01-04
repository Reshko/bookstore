package controllers

import (
	"net/http"

	"bookstore/models"
	"github.com/gin-gonic/gin"
)

func GetBooks(contex *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	contex.JSON(http.StatusOK, gin.H{"data": books})
}
