package models
import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
// SetupDB : initializing mysql database

func SetupDB() *gorm.DB {
	erro := godotenv.Load()
	if erro != nil {
		fmt.Println(erro)
	}
	var MYSQL_HOST = os.Getenv("MYSQL_HOST")
	var MYSQL_USER = os.Getenv("MYSQL_USER")
	var MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	var MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	var MYSQL_PORT = os.Getenv("MYSQL_PORT")
	
	if MYSQL_PORT == "" {
		MYSQL_PORT = "3306"
	}

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DBNAME)
	
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})
	
	if err != nil {
	panic(err.Error())
	}

	return db
}