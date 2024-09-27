package main

import (
	"github.com/gin-gonic/gin"
	"Backend/routers"
)

func main(){
    router := gin.Default()

	routers.UserRouter(router)

	router.Run(":9000")
	
}
