package app

import (
	"book-store/internal/conf"
	"book-store/internal/controllers"
)

func mapurls(appConf *conf.AppConfig) {
	bcontroller := controllers.NewBookController(appConf)
	api := router.Group("/book-store")
	api.POST("/")

}
