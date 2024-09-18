package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/docs"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/handler"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/users"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Setup database connection using GORM with PostgreSQL
	dsn := "host=127.0.0.1 user=postgres password=testing_password dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error()) // Log and terminate if the database connection fails
	}

	// Initialize repositories, services, and handlers for user-related operations
	// Repository layer to interact with database
	userRepository := users.UserRepository(db)
	// Service layer for business logic
	userService := users.UserServices(userRepository)
	// Handler for user-related HTTP requests
	userHandler := handler.UserHandler(userService)

	// Setup Gin router and API route groups
	router := gin.Default()
	api := router.Group("/api/v1")

	// Create endpoint for user registration
	api.POST("/users", userHandler.Signup)

	// Swagger API docs route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://127.0.0.1:8888/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// Start the server on localhost:9999
	router.Run("127.0.0.1:8888")
}
