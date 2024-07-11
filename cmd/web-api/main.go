package main

import (
	"interview/pkg/config"
	"interview/pkg/controllers"
	"interview/pkg/db"
	"interview/pkg/middleware"
	repo "interview/pkg/repository"
	services "interview/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	// Initialize database
	database := db.GetDatabase(cfg)

	// Initialize repository
	cartRepository := repo.NewRepository(database)

	// Initialize services
	cartService := services.NewService(cartRepository)

	// Initialize controller
	cartController := controllers.NewCartController(cartService)

	r := gin.Default()

	// Apply middleware
	r.Use(middleware.SessionMiddleware())

	// Define routes
	r.GET("/", cartController.ShowAddItemForm)
	r.POST("/add-item", cartController.AddItem)
	r.GET("/remove-cart-item", cartController.DeleteCartItem)

	// Run the server
	r.Run(":8088")
}
