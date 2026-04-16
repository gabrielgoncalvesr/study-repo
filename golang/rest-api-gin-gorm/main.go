package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gabrielgoncalvesr/study-repo/controllers"
	"github.com/gabrielgoncalvesr/study-repo/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
