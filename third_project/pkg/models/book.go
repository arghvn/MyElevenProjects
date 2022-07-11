package models

import (
	"github.com/arghvn/go.bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// for changing database create below function

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// What is difference between Create and New Record?
// NewRecord doesn't affect the database at all,
// it just returns true if the current value's primary key is unset, meaning it's a new record.
// This means calling NewRecord in void context, as you have done,
// is meaningless, since you're ignoring the return value.

func GetAllBooks() []Book {
	var Books []book
	db.Find(&Books)
	return Books
}

func GetBookId(Id int64) (*Book, *gorm.DB) {
	var getBook book
	db := db.where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var Book book
	db.Where("ID=?", ID).Delete(book)
	return book
}
