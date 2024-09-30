package routers

import (
	"Backend/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	{
		router.PUT("/update", userHandler.UpdateUserHandler)
		router.GET("/me", userHandler.GetUserDetailHandler)
	}
}
