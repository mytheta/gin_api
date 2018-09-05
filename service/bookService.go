package service

import (
	"github.com/mytheta/gin_api/model"
	"github.com/mytheta/gin_api/repository"
)

var Book = bookimpl{}

type bookimpl struct {
}

func (b *bookimpl) Create(book model.Book) error{
	return repository.Book.Create(book)
}
