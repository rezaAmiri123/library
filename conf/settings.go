package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}

var SecretKey string
func Init() {
	SecretKey = os.Getenv("SecretKey")
	DBULR := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("db_user"),
		os.Getenv("db_password"),
		os.Getenv("db_host"),
		os.Getenv("db_port"),
		os.Getenv("db_name"),
	)
	db, err := gorm.Open("mysql", DBULR)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}
	DB = db
}
