package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/items", controller.GetItem())
	incomingRoutes.GET("/items/:item_id", controller.GetItem())
	incomingRoutes.GET("/items-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/item", controller.CreateItem())
	incomingRoutes.PATCH("/items/:item_id", controller.UpdateItem())

}
