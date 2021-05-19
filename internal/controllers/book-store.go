package controllers

import (
	"book-store/internal/conf"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	AppConf *conf.AppConfig
}

func NewBookController(appconf *conf.AppConfig) *bookController {
	return &bookController{
		AppConf: appconf,
	}
}

func (bc *bookController) AddBook(c *gin.Context) {

}
