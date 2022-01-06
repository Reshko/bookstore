package main

import (
	"bookstore/cmd/cmd/controllers"
	_ "bookstore/cmd/cmd/docs"
	"bookstore/cmd/cmd/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

func main() {

	server := gin.Default()

	models.ConnectDatabase()

	v1 := server.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("", controllers.GetBooks)
			books.GET(":id", controllers.GetBookByID)
			books.POST("", controllers.PostBook)
			books.DELETE(":id", controllers.DeleteBookByID)
			books.PATCH(":id", controllers.PatchBookByID)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run("localhost:8080")
}
