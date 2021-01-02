package main

import (
	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/library/conf"
	"log"
)


func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := conf.InitDatabase()
	defer db.Close()
}
