package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/library/apps/accounts"
	"github.com/rezaAmiri123/library/conf"
	"log"
)

func RegisterApps(db *gorm.DB){
	accounts.AutoMigrate(db)
}
func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := conf.InitDatabase()
	RegisterApps(db)
	defer db.Close()
}
