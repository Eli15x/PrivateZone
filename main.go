package main

import (
	//kafka "ZCOM/src/client/kafka"
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	"PrivateZone/src/config"
	"PrivateZone/src/handler"
	"PrivateZone/src/service"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Carregar configurações
	config.LoadConfig()

	// Configuração dos serviços
	jwtSecretKey := config.GetConfig().JWTSecretKey
	authService := service.NewAuthService(jwtSecretKey)

	// Configuração dos controllers
	authController := handler.NewAuthController(authService)

	defer cancel()
	//Connection to Mongo
	/*if err := client.GetInstance().Initialize(ctx); err != nil {
		fmt.Errorf("mongo off")
		bugsnag.Notify(fmt.Errorf("[MONGO DB - Private Zone] Could not resolve Data access layer. Error:"))
	}


	/*if err := kafka.GetInstanceKafka().Initialize(); err != nil {
		fmt.Println(err)
		fmt.Errorf("Error initialize kafka Producer")
	}*/

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/login", authController.ValidateUser)
	//router.POST("/loginViaCode", handlers.ValidateUserViaCode)
	router.POST("/cadastro", authController.CreateUser)
	router.PUT("/user/edit", authController.EditUser)
	router.DELETE("/user/delete", authController.DeleteUser)
	router.GET("/user/getByUserId", authController.GetInformationByUserId)
	router.GET("/user/getByName", authController.GetUserByName)
	router.GET("/user/editImage", authController.GetUserByName)
	router.GET("/user/getByAcess", authController.GetUsersByAcess)
	router.GET("/users", authController.GetUsers)

	router.POST("/Content/File/create", authController.CreateProduct)
	router.POST("/Content/File/GetByContentId", authController.CreateProduct)
	router.POST("/Content/File/getByName", authController.CreateProduct)
	router.POST("/Content/File/delete", authController.CreateProduct)
	router.POST("/Content/File/edit", authController.CreateProduct)
	router.GET("/Content/Stream/create", authController.GetProductByName)
	router.GET("/Content/Stream/getByContentId", authController.GetProductByName)
	router.GET("/Content/Stream/edit", authController.GetCodeCest)
	router.GET("/Content/Stream/delete", authController.GetProducts)

	router.Run(":1323")
}
