package controllers

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	"book-store/internal/models"
	"book-store/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService *services.BookService
	AppConf     *conf.AppConfig
}

func NewBookController(appconf *conf.AppConfig) *BookController {
	return &BookController{
		BookService: services.NewBookService(appconf),
		AppConf:     appconf,
	}
}

func (bc *BookController) AddBook(c *gin.Context) {
	var book *models.Book

	//m := "Adding book"
	if err := c.ShouldBindJSON(&book); err != nil {
		fmt.Print("error while binding: ", err)
		return

	}

	if err := bc.BookService.AddBook(book); errors.HasError(&err) {
		fmt.Print("error while passing to servcie layer: ", err)
		return
	}

	c.JSON(http.StatusOK, book)
}
