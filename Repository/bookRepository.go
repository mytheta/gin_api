package repository

import (
	"github.com/mytheta/gin_api/model"
)
var Book = bookimpl{}

type bookimpl struct {
}

func (b *bookimpl) Create(book model.Book) error {
	err := db.Create(&book).Error
	return err
}

func (b *bookimpl) FindAll() ([]model.Book,error) {
	var books []model.Book
	err := db.Find(&books).Error
	return books, err
}