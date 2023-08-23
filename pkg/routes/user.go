package routes

import (
	"Teeverse/pkg/api/handler"
	"Teeverse/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, wishlistHandler *handler.WishlistHandler) {
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)
	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{

		payment := engine.Group("/payment")
		{
			payment.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
			payment.GET("/update_status", paymentHandler.VerifyPayment)
		}

		search := engine.Group("/search")
		{
			search.POST("/", inventoryHandler.SearchProducts)
		}

		home := engine.Group("/home")
		{
			home.GET("/products", inventoryHandler.ListProducts)
			home.GET("/products/details", inventoryHandler.ShowIndividualProducts)
			home.POST("/add-to-cart", cartHandler.AddToCart)
			home.POST("/add-to-wishlist", wishlistHandler.AddToWishlist)

		}
		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)

			edit := profile.Group("/edit")
			{
				edit.PUT("/name", userHandler.EditName)
				edit.PUT("/email", userHandler.EditEmail)
				edit.PUT("/phone", userHandler.EditPhone)
			}

			security := profile.Group("/security")
			{
				security.PUT("/change-password", userHandler.ChangePassword)
			}
			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.DELETE("", orderHandler.CancelOrder)

			}
		}

		cart := engine.Group("/cart")
		{
			cart.GET("/", userHandler.GetCart)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
			cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
			cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)

		}
		wishlist := engine.Group("/wishlist")
		{
			wishlist.GET("/", wishlistHandler.GetWishlist)
			wishlist.DELETE("/remove", wishlistHandler.RemoveFromWishlist)

		}

		checkout := engine.Group("/check-out")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}

	}
}
