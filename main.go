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
	operationTypeController := infra.SetupOperationTypeController()
	transactionController := infra.SetupTransactionController()

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	account := r.Group("/accounts")
	account.GET("", accountController.ListAll)
	account.POST("", accountController.Create)
	account.GET("/:id", accountController.Get)
	account.PUT("/:id", accountController.Update)
	account.DELETE("/:id", accountController.Delete)

	operationType := r.Group("/operation-types")
	operationType.GET("", operationTypeController.ListAll)
	operationType.POST("", operationTypeController.Create)
	operationType.GET("/:id", operationTypeController.Get)
	operationType.PUT("/:id", operationTypeController.Update)
	operationType.DELETE("/:id", operationTypeController.Delete)

	transaction := r.Group("/transactions")
	transaction.GET("/", transactionController.ListAll)
	transaction.POST("", transactionController.Create)
	transaction.GET("/:id", transactionController.Get)
	transaction.PUT("/:id", transactionController.Update)
	transaction.DELETE("/:id", transactionController.Delete)

	r.Run() // localhost:8080
}
