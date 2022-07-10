package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// dor connetction to database

func Connect() {
	d, err := gorm.Open("mysql", "arghvn:ABCD/simplerest?charset=utf8&parsetime=True&loc=local")
	if err != nil {
		panic(err)
	}
	db = d
}

// arghvn is username
// ABCD is password
// simplerest is name of my table
// database requirement

func GetDB() *gorm.DB {
	return db
}
