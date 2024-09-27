package routers

import (
    "github.com/gin-gonic/gin"
    "Backend/controllers"
)

func UserRouter(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/signup", controllers.SignUp)
        userRoutes.POST("/login", controllers.Login)
        userRoutes.PUT("/:user_id", controllers.UpdateUser)
        userRoutes.GET("/:user_id", controllers.GetUser)
        userRoutes.DELETE("/:user_id", controllers.DeleteUser)
    }
}
