package main

import (
	"interview/pkg/controllers"
	"interview/pkg/db"
	repo "interview/pkg/repository"
	service "interview/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.GetDatabase()

	ginEngine := gin.Default()

	cartRepository := repo.NewCartRepository(database)
	cartService := service.NewCartService(cartRepository)

	cartController := controllers.NewCartController(cartService)

	ginEngine.GET("/", cartController.ShowAddItemForm)
	ginEngine.POST("/add-item", cartController.AddItem)
	ginEngine.GET("/remove-cart-item", cartController.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
