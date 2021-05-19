package app

import (
	"book-store/internal/conf"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication(appConf *conf.AppConfig) {
	router.Use(gin.Recovery())
	mapurls(appConf)

	log.Printf("Starting service: %v on port %v\n", appConf.Server.APINAME, appConf.Server.APIPORT)

	router.Run(fmt.Sprintf(":%v", appConf.Server.APIPORT))
}
