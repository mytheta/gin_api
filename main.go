package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/controller"
)

func main() {
	r := gin.Default()
	r.POST("/book", controller.Book.Create)
	r.GET("/book", controller.Book.FindAll)
	r.PUT("/book/@:id",controller.Book.Update)
	r.DELETE("/book/@:id",controller.Book.Delete)
	r.Run(":8000")
}
