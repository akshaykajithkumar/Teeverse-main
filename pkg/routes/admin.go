package routes

import (
	"Teeverse/pkg/api/handler"
	"Teeverse/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler /*userHandler *handler.UserHandler,*/, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, offerHandler *handler.OfferHandler, couponHandler *handler.CouponHandler) {
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
			categorymanagement.GET("/", categoryHandler.Categories)
			categorymanagement.POST("/add", categoryHandler.AddCategory)
			categorymanagement.PUT("/update", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		inventorymanagement := engine.Group("/inventories")
		{
			//inventorymanagement.GET("", inventoryHandler.ListProducts)
			//inventorymanagement.GET("/details", inventoryHandler.ShowIndividualProducts)
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
		sales := engine.Group("/sales")
		{
			sales.GET("/daily", orderHandler.AdminSalesDailyReport)
			sales.GET("/weekly", orderHandler.AdminSalesWeeklyReport)
			sales.GET("/monthly", orderHandler.AdminSalesMonthlyReport)
			sales.GET("/annual", orderHandler.AdminSalesAnnualReport)
			sales.POST("/custom", orderHandler.AdminSalesCustomReport)
		}
		products := engine.Group("/products")
		{
			products.GET("", inventoryHandler.AdminListProducts)
			products.GET("/details", inventoryHandler.ShowIndividualProducts)
			products.GET("/search", inventoryHandler.SearchProducts)
			products.GET("/category", inventoryHandler.GetCategoryProducts)
		}
		offers := engine.Group("/offers")
		{
			offers.POST("/create", offerHandler.AddOffer)
			offers.POST("/expire", offerHandler.ExpireValidity)
		}

		coupons := engine.Group("/coupons")
		{
			coupons.POST("/create", couponHandler.CreateNewCoupon)
			coupons.POST("/expire", couponHandler.MakeCouponInvalid)
		}
	}
}
