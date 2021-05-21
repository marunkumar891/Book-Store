package controllers

import (
	"book-store/internal/conf"
	"book-store/internal/errors"
	"book-store/internal/models"
	"book-store/internal/services"
	"log"
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
		log.Print("error while binding: ", err)
		restErr := errors.NewBadRequestError("invalid request body: " + err.Error())
		c.JSON(restErr.Status, restErr)
		return

	}

	if err := bc.BookService.AddBook(book); errors.HasError(&err) {
		log.Print("error while passing to servcie layer: ", err)
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, book)
}
func (bc *BookController) GetAllBook(c *gin.Context) {

	books, err := bc.BookService.GetAllBook()
	if errors.HasError(&err) {
		log.Print("error while passing to service layer: ", err)
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, books)
}
func (bc *BookController) GetBook(c *gin.Context) {
	book := models.Book{}
	id := c.Params.ByName("ID")
	if id == "" {
		log.Print("id invalid")
		return
	}
	book, err := bc.BookService.GetBook(book, id)
	if errors.HasError(&err) {
		log.Print("error while passing to service layer: ", err)
		c.JSON(err.Status, err)
		return
	}
	c.JSON(err.Status, book)

}
func (bc *BookController) CheckBook(c *gin.Context) {
	book := models.Book{}
	id := c.Params.ByName("ID")
	if id == "" {
		log.Print("id invalid")
		return
	}
	book, err := bc.BookService.CheckBook(book, id)
	if errors.HasError(&err) {
		log.Print("error while passing to service layer: ", err)
		c.JSON(err.Status, err)
		return
	}
	if book.IsAvailable {
		c.JSON(err.Status, "Book available")

	} else {
		c.JSON(err.Status, "Book unavialable")

	}

}
