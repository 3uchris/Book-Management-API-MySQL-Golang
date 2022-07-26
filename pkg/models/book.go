package models

import (
	"github.com/3uchris/Book-Management-API-MySQL-Golang/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string 'gorm:"" json:"name"'
	Author string 'json:"author"'
	Publication string 'json:"publication"'
}

func init(){
	config.Connect()
	db = config.GetDB() //from app.go
	db.AutoMigrate(&Book{})
}
