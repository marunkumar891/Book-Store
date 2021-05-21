package services

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	"book-store/internal/models"
	"book-store/internal/repo"
	"log"
	"time"
)

type BookService struct {
	bookrepo *repo.BookRepo
	AppConf  *conf.AppConfig
}

func NewBookService(appconf *conf.AppConfig) *BookService {
	return &BookService{
		AppConf:  appconf,
		bookrepo: repo.NewBookRepo(appconf),
	}
}

func (bs *BookService) AddBook(book *models.Book) errors.RestAPIError {
	book.AddedAt = time.Now()
	if err := bs.bookrepo.AddBook(book); errors.HasError(&err) {
		log.Print("error while passing to repo layer: ", err)
		return err
	}
	return errors.NO_ERROR()
}
func (bs *BookService) GetAllBook() ([]models.Book, errors.RestAPIError) {
	books := []models.Book{}
	err := bs.bookrepo.GetAllBook(&books)
	if errors.HasError(&err) {
		log.Print("error while passing to repo layer: ", err)
		return []models.Book{}, err
	}

	return books, errors.NO_ERROR()
}
func (bs *BookService) GetBook(book models.Book, id string) (models.Book, errors.RestAPIError) {
	book = models.Book{}
	err := bs.bookrepo.GetBook(&book, id)
	if errors.HasError(&err) {
		log.Print("error while passing to repo layer: ", err.Message)
		return models.Book{}, err
	}
	return book, errors.NO_ERROR()

}
func (bs *BookService) CheckBook(book models.Book, id string) (models.Book, errors.RestAPIError) {
	book = models.Book{}
	err := bs.bookrepo.CheckBook(&book, id)
	if errors.HasError(&err) {
		log.Print("error while passing to repo layer: ", err.Message)
		return models.Book{}, err
	}
	return book, errors.NO_ERROR()

}
