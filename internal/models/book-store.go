package models

import "time"

const Table_Book = "db_book"

type Book struct {
	Id      uint64    `json:"id"`
	Author  string    `json:"author"`
	AddedAt time.Time `json:"create_at"`
}
