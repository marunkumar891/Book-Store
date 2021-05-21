package app

import (
	"book-store/internal/conf"
	"book-store/internal/controllers"
)

func mapurls(appConf *conf.AppConfig) {
	bcontroller := controllers.NewBookController(appConf)

	api := router.Group("/bookstore")
	{
		api.POST("/add/book", bcontroller.AddBook)
		api.GET("/get/all/book", bcontroller.GetAllBook)
		api.GET("/get/book/:ID", bcontroller.GetBook)
		api.GET("/check/book/:ID", bcontroller.CheckBook)
		api.DELETE("/delete/book/:ID", bcontroller.DeleteBook)

	}

}
