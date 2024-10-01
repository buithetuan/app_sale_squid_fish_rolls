package main

import (
	"Backend/logs"
	"Backend/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	logs.InitLogger()

	router := gin.Default()

	routers.InitRouter(router)

	router.Run(":9000")

}
