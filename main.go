package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Connection GORM to database
	dsn := "host=127.0.0.1 user=postgres password=testing_password dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	c := gin.Default()

	//Users
	userRepository := users.UserRepository(db)

	c.Run("127.0.0.1:8080")
}
