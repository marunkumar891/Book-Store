package models

import (
	"time"
)

const Table_Book = "books"

type Book struct {
	//gorm.Model
	Id          string    `json:"id" gorm:"column:id;primaryKey"`
	BookName    string    `json:"book_name" gorm:"column:book_name"`
	Author      string    `json:"author" gorm:"column:author"`
	IsAvailable bool      `json:"is_available" gorm:"column:is_available"`
	AddedAt     time.Time `json:"added_at" gorm:"column:added_at"`
}
