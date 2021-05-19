package models

import "time"

const Table_Book = "books"

type Book struct {
	Id      uint64    `json:"id" gorm:"primary_key"`
	Author  string    `json:"author"`
	AddedAt time.Time `json:"added_at"`
}
