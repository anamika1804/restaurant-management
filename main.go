package main

import (
	"golang-restaurant-management/database"
	"golang-restaurant-management/middleware"
	"golang-restaurant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.ItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}
