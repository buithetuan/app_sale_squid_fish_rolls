package routers

import (
	"Backend/config"
	"Backend/handlers"
	"Backend/repo"
	"Backend/services"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	cnf := config.LoadConfig()
	gormDB, err := config.ConnectDB(cnf)
	if err != nil {
		panic(err)
	}
	userRepo := repo.NewUserRepo(gormDB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	v1 := router.Group("/api/v1")
	UserRouter(v1, userHandler)
}
