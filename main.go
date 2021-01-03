package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rezaAmiri123/library/apps/accounts"
	"github.com/rezaAmiri123/library/conf"
	"log"
)

func RegisterApps(rg *gin.RouterGroup){
	accounts.Register(rg)
}
func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := conf.InitDatabase()
	defer db.Close()
	router := gin.Default()
	v1 := router.Group("/api")
	RegisterApps(v1)
	router.Run()
}
