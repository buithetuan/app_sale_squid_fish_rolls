package routers

import (
	"Backend/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup, authHandler *handlers.AuthHandler) {
	{
		router.POST("/signup", authHandler.SignUpHandler)
		router.POST("/login", authHandler.LogInHandler)
		router.POST("/refresh-token", authHandler.RefreshTokenHandler)
	}
}
