package main

import (
	"interview/pkg/controllers"
	"interview/pkg/db"
	"interview/pkg/middleware"
	repo "interview/pkg/repository"
	services "interview/pkg/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database
	database := db.GetDatabase()

	// Initialize repository
	cartRepository := repo.NewCartRepository(database)

	// Initialize services
	cartService := services.NewCartService(cartRepository)

	// Initialize controller
	cartController := controllers.NewCartController(cartService)

	// Apply middleware
	r.Use(middleware.SessionMiddleware())

	// Define routes
	r.GET("/", cartController.ShowAddItemForm)
	r.POST("/add-item", cartController.AddItem)
	r.GET("/remove-cart-item", cartController.DeleteCartItem)

	// Run the server
	r.Run(":8088")
}
