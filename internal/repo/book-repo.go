package repo

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	"book-store/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type BookRepo struct {
	DB      *gorm.DB
	AppConf *conf.AppConfig
}

func NewBookRepo(appconf *conf.AppConfig) *BookRepo {
	return &BookRepo{
		DB:      appconf.DB,
		AppConf: appconf,
	}
}

func (br *BookRepo) AddBook(book *models.Book) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Create(book).Error; err != nil {
		log.Print("unable to add book into db: ", err.Error())
		return errors.NewInternalServerError(err.Error())
	}
	return errors.NO_ERROR()
}
func (br *BookRepo) GetAllBook(books *[]models.Book) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Find(books).Error; err != nil {
		log.Print("unable to fetch from db: ", err.Error())
		return errors.NewInternalServerError(err.Error())
	}
	return errors.NO_ERROR()
}
func (br *BookRepo) GetBook(book *models.Book, id string) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Where("id = ?", id).Find(book).Error; err != nil {
		log.Print("unable to fetch from db: ", err.Error())
		return errors.NewNotFoundError(err.Error())
	}
	return errors.NO_ERROR()
}
func (br *BookRepo) CheckBook(book *models.Book, id string) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Where("id = ?", id).Find(book).Error; err != nil {
		log.Print("unable to fetch from db: ", err.Error())
		return errors.NewNotFoundError(err.Error())
	}
	return errors.NO_ERROR()
}

func (br *BookRepo) DeleteBook(id string, book *models.Book) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Where("id = ?", id).Delete(&book).Error; err != nil {
		log.Print("unable to delete record from db: ", err)
		return errors.NewInternalServerError(err.Error())
	}
	//br.DB.Delete(&book)
	return errors.NO_ERROR()
}
