package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const Table_Book = "books"

type Book struct {
	gorm.Model
	Id          uint64    `json:"id" gorm:"column:id;primary_key;unique"`
	BookName    string    `json:"book_name" gorm:"column:book_name"`
	Author      string    `json:"author" gorm:"column:author"`
	IsAvailable bool      `json:"is_available" gorm:"column:is_available"`
	AddedAt     time.Time `json:"added_at" gorm:"column:added_at"`
}
