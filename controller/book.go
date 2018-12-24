package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/service"
	"log"
	"net/http"
	"strconv"
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
		"message": "本登録成功",
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

func (b *bookimpl) Update(c *gin.Context){
	var book model.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	err = service.Book.Update(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"編集できました,"})

}

func (b *bookimpl) Delete(c *gin.Context){
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	ok,err := service.Book.IsExistByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,nil)
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound,nil)
	}

	err = service.Book.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"削除成功"})
	}