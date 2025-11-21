package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rigle/rigle-go/config"
	"github.com/rigle/rigle-go/handlers"
)

func main() {
	// Initialize database
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create handler
	itemHandler := handlers.NewItemHandler(db)

	// Setup router
	router := gin.Default()

	// Item routes
	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.GET("", itemHandler.GetAllItems)
			items.GET("/:id", itemHandler.GetItem)
			items.POST("", itemHandler.CreateItem)
			items.PUT("/:id", itemHandler.UpdateItem)
			items.DELETE("/:id", itemHandler.DeleteItem)
		}
	}

	// Start server
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
