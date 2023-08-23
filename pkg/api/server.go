package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "Teeverse/cmd/api/docs"
	handler "Teeverse/pkg/api/handler"
	"Teeverse/pkg/routes"
)

// ServerHTTP represents an HTTP server for the web application.
type ServerHTTP struct {
	engine *gin.Engine // engine is the core of the Gin web framework, responsible for routing HTTP requests and handling middleware.
}

/*
NewServerHTTP creates a new instance of ServerHTTP.

Parameters:
- categoryHandler: A handler for category-related operations.
- inventoryHandler: A handler for inventory-related operations.
- userHandler: A handler for user-related operations.
- otpHandler: A handler for OTP-related operations.
- adminHandler: A handler for admin-related operations.
- cartHandler: A handler for cart-related operations.
- orderHandler: A handler for order-related operations.

Returns:
- *ServerHTTP: A pointer to the newly created ServerHTTP instance.
*/
func NewServerHTTP(categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, adminHandler *handler.AdminHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, wishlistHandler *handler.WishlistHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("pkg/templates/*.html")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.UserRoutes(engine.Group("/users"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, wishlistHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, userHandler, categoryHandler, inventoryHandler, orderHandler, paymentHandler)

	return &ServerHTTP{engine: engine}
}

/*
Start starts the HTTP server and listens on port 1243.
*/
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":1243")
}
