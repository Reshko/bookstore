package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{"data": "hello golang"})
	})

	server.Run("localhost:8080")
}
