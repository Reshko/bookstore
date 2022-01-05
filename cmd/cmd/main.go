package main

import (
	"bookstore/cmd/cmd/controllers"
	"bookstore/cmd/cmd/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	models.ConnectDatabase()

	server.GET("/books", controllers.GetBooks)
	server.POST("/books", controllers.PostBook)
	server.GET("/books/:id", controllers.GetBookByID)
	server.DELETE("/books/:id", controllers.DeleteBookByID)
	server.PATCH("/books/:id", controllers.PatchBookByID)

	server.Run("localhost:8080")
}
