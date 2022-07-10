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
