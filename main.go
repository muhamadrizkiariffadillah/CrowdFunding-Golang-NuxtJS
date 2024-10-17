package main

import (
	"log"

	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/authJWT"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/campaigns"
	"github.com/muhamadrizkiariffadillah/CrowdFunding-Golang-NuxtJS/transaction"

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
	dsn := "host=127.0.0.1 user=postgres password=123qweasdzxc dbname=crowdfunding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error()) // Log and terminate if the database connection fails
	}
	authService := authJWT.NewJwtService()
	// Initialize repositories, services, and handlers for user-related operations
	// Repository layer to interact with database
	userRepository := users.UserRepository(db)
	// Service layer for business logic
	userService := users.UserServices(userRepository)
	// Handler for user-related HTTP requests
	userHandler := handler.UserHandler(userService, authService)
	//
	middleware := authJWT.MiddleWare(authService, userService)

	// Campaign
	campaignRepository := campaigns.CampaignRepository(db)
	// Service layer for business logic
	campaignServices := campaigns.CampaignServices(campaignRepository)
	// Handler for compaign-related HTTP requests.
	campaignHandler := handler.CampaignHandler(campaignServices)

	// transaction
	transactionRepository := transaction.TransactionsRepository(db)
	transactionServices := transaction.TransactionsServices(transactionRepository)
	transactionHandler := handler.TransactionsHandler(transactionServices)

	// Setup Gin router and API route groups
	router := gin.Default()
	// Static router avatar
	router.Static("/images/user/avatar", "./images")
	
	api := router.Group("/api/v1")

	// user url
	//  Endpoint for user registration
	api.POST("/users/signup", userHandler.Signup)
	// Endpoint for user logged in
	api.POST("/users/login", userHandler.Login)
	// Endpoint for fetching the user.
	api.GET("/users/me", userHandler.FetchUser)
	// Endpoint for checking email user is available.
	api.POST("/users/check-email", userHandler.CheckEmail)
	// Endpoint for uploading user avatar
	api.POST("/users/me/upload-avatar", middleware, userHandler.UploadAvatar)

	// campaign url
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaign", campaignHandler.GetCampaign)
	api.POST(("/campaign/create"), middleware, campaignHandler.CreateCampaign)
	api.POST("/campaign/update", middleware, campaignHandler.UpdateCampaign)

	// campaign images
	api.POST("campaign/image", middleware, campaignHandler.SaveCampaignImage)

	// transaction api
	api.GET("/campaign/transactions", transactionHandler.GetCampaignTransactions)
	api.GET("/user/transactions", middleware, transactionHandler.GetUserTransactions)

	// Swagger API docs route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("http://127.0.0.1:8888/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// Start the server on localhost:9999
	router.Run("127.0.0.1:8888")
}
