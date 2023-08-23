package routes

import (
	"Teeverse/pkg/api/handler"
	"Teeverse/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, userHandler *handler.UserHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		usermanagement := engine.Group("/users")
		{
			usermanagement.POST("/block", adminHandler.BlockUser)
			usermanagement.POST("/unblock", adminHandler.UnBlockUser)
			usermanagement.GET("/getusers", adminHandler.GetUsers)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("/add", inventoryHandler.AddInventory)
			inventorymanagement.PUT("/update", inventoryHandler.UpdateInventory)
			inventorymanagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}
		orders := engine.Group("/orders")
		{
			orders.PUT("/edit/status", orderHandler.EditOrderStatus)
			orders.GET("", orderHandler.AdminOrders)
		}
		paymentmethods := engine.Group("/paymentmethods")
		{
			paymentmethods.GET("/", paymentHandler.GetPaymentMethods)
			paymentmethods.POST("/add", paymentHandler.AddNewPaymentMethod)
			paymentmethods.DELETE("/remove", paymentHandler.RemovePaymentMethod)
		}

	}
}
