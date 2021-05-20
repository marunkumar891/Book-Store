package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const Table_Book = "books"

type Book struct {
	gorm.Model
	Id      uint64    `json:"id" gorm:"primary_key"`
	Author  string    `json:"author"`
	AddedAt time.Time `json:"added_at"`
}
