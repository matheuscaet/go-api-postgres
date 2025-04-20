package main

import (
	controller "go-api-postgres/controllers"
	"go-api-postgres/db"
	"go-api-postgres/repository"
	"go-api-postgres/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ItemRepository := repository.NewItemRepository(dbConnection)
	ItemUseCase := usecase.NewItemUseCase(ItemRepository)
	ItemController := controller.NewItemController(ItemUseCase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/items", ItemController.GetItems)
	server.POST("/item", ItemController.CreateItem)
	server.GET("/item/:itemId", ItemController.GetItemById)

	server.Run(":" + os.Getenv("APP_PORT"))

}
