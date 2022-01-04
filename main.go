package main

import (
	"github.com/gin-gonic/gin"

	"bookstore/controllers"
	"bookstore/models"
)

func main() {
	server := gin.Default()

	models.ConnectDatabase()

	server.GET("/books", controllers.GetBooks)

	server.Run("localhost:8080")
}