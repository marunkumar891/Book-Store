package repo

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	"book-store/internal/models"

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
