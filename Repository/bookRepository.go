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

func (b *bookimpl) Update(book model.Book) error {
	err := db.Save(&book).Error
	return err
}

func (b *bookimpl) Delete(id uint) error {
	book := model.Book{}
	err :=  db.Delete(&book,id).Error
	return err
}

func (b *bookimpl) IsExistByID(id uint) (bool,error){
	var books []model.Book
	err := db.Where("id = ?",id).Find(&books).Error
	return len(books) >= 1,err
}