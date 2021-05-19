package main

import (
	"book-store/internal/app"
	"book-store/internal/conf"
	"book-store/internal/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var appConfig *conf.AppConfig

func init() {
	appConfig = conf.NewConfig()
}
func main() {
	fmt.Println("connecting to application")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&charset=utf8&autocommit=false",
		appConfig.DBConfig.DBUSER,
		appConfig.DBConfig.DBPASSWORD,
		appConfig.DBConfig.DBHOST,
		appConfig.DBConfig.DBPORT,
		appConfig.DBConfig.DBNAME,
	))
	//fmt.Println("cam here")
	if err != nil {
		log.Println("Error Connecting to the DB", err.Error())
		panic(err)
	}
	appConfig.DB = db
	if err := db.AutoMigrate(&models.Book{}); err.Error != nil {
		log.Print(err)
		fmt.Println("unable to auto migrate.")
	}
	//fmt.Print("came here")
	defer db.Close()
	app.StartApplication(appConfig)

}
