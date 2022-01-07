package controllers

import (
	"bookstore/cmd/cmd/httputil"
	"bookstore/cmd/cmd/models"
	"github.com/gin-gonic/gin"
	"net/http"
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

// GetBookByID godoc
// @Summary      Show a book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Book id"
// @Success      200  {object}  model.Book
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func GetBookByID(contex *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", contex.Param("id")).First(&book).Error; err != nil {
		//ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		httputil.NewError(contex, http.StatusNotFound, err)
		return
	}

	contex.JSON(http.StatusOK, gin.H{"data": book})

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

func PatchBookByID(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(models.Book{Title: input.Title, Author: input.Author})

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}
