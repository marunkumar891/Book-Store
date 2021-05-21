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
	if err := br.DB.Table(models.Table_Book).Save(book).Error; err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return errors.NO_ERROR()
}
func (br *BookRepo) GetAllBook(books *[]models.Book) errors.RestAPIError {
	if err := br.DB.Table(models.Table_Book).Find(books).Error; err != nil {
		log.Print("unable to fetch from db: ", err)
		return errors.NewInternalServerError(err.Error())
	}
	return errors.NO_ERROR()
}
