package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rezaAmiri123/library/apps/accounts"
	"log"
	"os"
)

var DB *gorm.DB

type Database struct {
	*gorm.DB
}

func RegisterApps(db *gorm.DB){
	accounts.AutoMigrate(db)
}
// Opening a database and save the reference to `Database` struct.
func InitDatabase() *gorm.DB {
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
	RegisterApps(db)
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}