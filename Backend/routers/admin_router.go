package routers

import (
	"Backend/handlers"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.RouterGroup, adminHandler *handlers.AdminHandler) {
	{
		router.POST("/add-product", adminHandler.AddProductHandler)
		router.DELETE("/remove-product", adminHandler.RemoveProductHandler)
		router.PUT("/update-product", adminHandler.UpdateProductStockHandler)
		router.PUT("/update-status-order", adminHandler.UpdateStatusOrdersHandler)
	}
}
