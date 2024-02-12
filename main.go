package main

import (
	kafka "ZCOM/src/client/kafka"
	//"time"
	//"context"

	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"ZCOM/src/client"
	"ZCOM/src/handler"
	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//bugsnag configure
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       "3ecac0ed23b7b1f4b863073135c602b8",
		ReleaseStage: "production",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
		// more configuration options
	})

	err := godotenv.Load(".env")
    if err != nil {
        fmt.Errorf("Error loading .env file")
    }

	//Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	//Connection to Mongo
	if err := client.GetInstance().Initialize(ctx); err != nil {
		fmt.Errorf("mongo off")
		bugsnag.Notify(fmt.Errorf("[MONGO DB - Private Zone] Could not resolve Data access layer. Error:"))
	}


	if err := kafka.GetInstanceKafka().Initialize(); err != nil {
		fmt.Println(err)
		fmt.Errorf("Error initialize kafka Producer")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.POST("/login", handlers.ValidateUser)
	//router.POST("/loginViaCode", handlers.ValidateUserViaCode)
	router.POST("/cadastro", handlers.CreateUser)
	router.PUT("/user/edit", handlers.EditUser)
	router.DELETE("/user/delete", handlers.DeleteUser)
	router.GET("/user/getByUserId", handlers.GetInformationByUserId)
	router.GET("/user/getByName", handlers.GetUserByName)
	router.GET("/user/editImage", handlers.GetUserByName)
	router.GET("/user/getByAcess", handlers.GetUsersByAcess)
	router.GET("/users", handlers.GetUsers)


	router.POST("/Content/File/create", handlers.CreateProduct)
	router.POST("/Content/File/GetByContentId", handlers.CreateProduct)
	router.POST("/Content/File/getByName", handlers.CreateProduct)
	router.POST("/Content/File/delete", handlers.CreateProduct)
	router.POST("/Content/File/edit", handlers.CreateProduct)
	router.GET("/Content/Stream/create", handlers.GetProductByName)
	router.GET("/Content/Stream/getByContentId", handlers.GetProductByName)
	router.GET("/Content/Stream/edit", handlers.GetCodeCest)
	router.GET("/Content/Stream/delete", handlers.GetProducts)



	router.Run(":1323")
}