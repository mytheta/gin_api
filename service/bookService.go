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

func (b *bookimpl) FindAll()  ([]model.Book,error) {
	return repository.Book.FindAll()
}

func (b *bookimpl) Update(book model.Book)  error {
	return repository.Book.Update(book)
}

func (b *bookimpl) Delete(id uint) error {
	return repository.Book.Delete(id)
}

func (b *bookimpl) IsExistByID(id uint) (bool,error) {
	return repository.Book.IsExistByID(id)
}