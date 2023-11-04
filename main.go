package main

import (
	"log/slog"

	"pismo/config"
	"pismo/docs"
	_ "pismo/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	r := gin.Default()

	infra := config.NewInfra()
	accountController := infra.SetupAccountController()

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/accounts", accountController.ListAllAccount)
	r.GET("/account/:id", accountController.GetAccount)
	r.POST("/account", accountController.CreateAccount)
	r.PUT("/account/:id", accountController.UpdateAccount)
	r.DELETE("/account/:id", accountController.DeleteAccount)

	r.Run() // localhost:8080
}
