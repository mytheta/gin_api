package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mytheta/gin_api/model"
)

type BookRepository interface {
	Create(book model.Book) error
	FindAll() ([]model.Book, error)
	Update(book model.Book) error
	Delete(id uint) error
	IsExistByID(id uint) (bool, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (b *bookRepository) Create(book model.Book) error {
	err := b.db.Create(&book).Error
	return err
}

func (b *bookRepository) FindAll() ([]model.Book, error) {
	var books []model.Book
	err := b.db.Find(&books).Error
	return books, err
}

func (b *bookRepository) Update(book model.Book) error {
	err := b.db.Save(&book).Error
	return err
}

func (b *bookRepository) Delete(id uint) error {
	book := model.Book{}
	err := b.db.Delete(&book, id).Error
	return err
}

func (b *bookRepository) IsExistByID(id uint) (bool, error) {
	var books []model.Book
	err := b.db.Where("id = ?", id).Find(&books).Error
	return len(books) >= 1, err
}
