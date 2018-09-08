package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/service"
	"log"
	"net/http"
)

var Book = bookimpl{}

type bookimpl struct {
}

func (b *bookimpl) Create(c *gin.Context) {

	var book model.Book
	err := c.BindJSON(&book)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = service.Book.Create(book)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ユーザー作成成功",
	})
}

func (b *bookimpl) FindAll(c *gin.Context){
	books, err := service.Book.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, books)
}