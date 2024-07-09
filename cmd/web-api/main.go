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

	var taxController controllers.TaxController
	ginEngine.GET("/", taxController.ShowAddItemForm)
	ginEngine.POST("/add-item", taxController.AddItem)
	ginEngine.GET("/remove-cart-item", taxController.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
