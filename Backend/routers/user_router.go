package routers

import (
	"Backend/handlers"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup, userHandler *handlers.UserHandler) {
	{
		router.POST("/users/signup", userHandler.SignUpHandler)
		router.POST("/login", userHandler.LogInHandler)
		router.PUT("users/update-user", userHandler.UpdateUserHandler)
		router.GET("/get-user-detail/:user_id", userHandler.GetUserDetailHandler)
	}
}
