package routes

import (
	"Teeverse/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(engine *gin.RouterGroup, inventoryHandler *handler.InventoryHandler) {
	engine.GET("", inventoryHandler.ListProducts)
	engine.GET("/details", inventoryHandler.ShowIndividualProducts)
	engine.GET("/search", inventoryHandler.SearchProducts)
	engine.GET("/category", inventoryHandler.GetCategoryProducts)
}
