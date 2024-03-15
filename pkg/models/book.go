package models

import (
	"github.com/pfirulo2022/go-bookstore.git/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `gorm:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	//db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	err := db.Find(&Books).Error
	if err != nil {
		return nil, err
	}
	return Books, nil

}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}
