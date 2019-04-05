package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/service"
)

type BookController interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(s service.BookService) BookController {
	return &bookController{bookService: s}
}

func (b *bookController) Create(c *gin.Context) {

	var book model.Book
	err := c.BindJSON(&book)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = b.bookService.Create(book)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "本登録成功",
	})
}

func (b *bookController) FindAll(c *gin.Context) {
	books, err := b.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b *bookController) Update(c *gin.Context) {
	var book model.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = b.bookService.Update(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "編集できました,"})

}

func (b *bookController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	ok, err := b.bookService.IsExistByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, nil)
	}

	err = b.bookService.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "削除成功"})
}
