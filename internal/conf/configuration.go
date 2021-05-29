package conf

import (
	//"log"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type dbConfig struct {
	DBUSER     string
	DBPASSWORD string
	DBHOST     string
	DBNAME     string
	DBPORT     string
	DBLOGMODE  bool
}
type loggingConfig struct {
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
}
type serverConfig struct {
	APIPORT      int
	APINAME      string
	ADMIN_KEY    string
	ADMIN_SECRET string
}

type AppConfig struct {
	Server        serverConfig
	DBConfig      dbConfig
	DB            *gorm.DB
	LoggingConfig loggingConfig
}

func NewConfig() *AppConfig {

	appConfig := AppConfig{

		Server: serverConfig{
			APIPORT:      getEnvAsInt("APIPORT", 8787),
			APINAME:      getEnv("APINAME", "book-store"),
			ADMIN_KEY:    getEnv("ADMIN_KEY", "admin"),
			ADMIN_SECRET: getEnv("ADMIN_SECRET", "admin@admin"),
		},

		DBConfig: dbConfig{
			DBUSER:     getEnv("DBUSER", "root"),
			DBPASSWORD: getEnv("DBPASSWORD", "Tarus_891"),
			DBNAME:     getEnv("DBNAME", "bookstore"),
			DBHOST:     getEnv("DBHOST", "localhost"),
			DBPORT:     getEnv("DBPORT", "3306"),
			DBLOGMODE:  getEnvAsBool("DBLOGMODE", false),
		},
	}
	return &appConfig
}

func getEnv(key string, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultval
}
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, strconv.Itoa(defaultVal))
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
