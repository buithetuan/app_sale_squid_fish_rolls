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
	productRepo := repo.NewProductRepo(gormDB)
	orderRepo := repo.NewOrderRepo(gormDB)
	cartRepo := repo.NewCartRepo(gormDB)
	cartItemsRepo := repo.NewCartItemRepo(gormDB)

	userService := services.NewUserService(userRepo, cartRepo, orderRepo, cartItemsRepo)
	authService := services.NewAuthService(userRepo, cartRepo)
	adminService := services.NewAdminService(userRepo, productRepo, orderRepo)

	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	adminHandler := handlers.NewAdminHandler(adminService)

	v1a := router.Group("/api/v1/auth")
	AuthRouter(v1a, authHandler)
	v1u := router.Group("/api/v1/users")
	UserRouter(v1u, userHandler)
	v1ad := router.Group("/api/v1/admin")
	AdminRouter(v1ad, adminHandler)
}
