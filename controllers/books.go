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

func PostBook(contex *gin.Context) {
	var input models.CreateBookInput
	var books []models.Book

	if err := contex.ShouldBindJSON(&input); err != nil {
		contex.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	models.DB.Find(&books)

	contex.JSON(http.StatusOK, gin.H{"data": books})
}

func GetBookByID(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})

}

func DeleteBookByID(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})

}
