package router

import "github.com/gin-gonic/gin"

func Init(addr string) {
	router := gin.Default()

	initializeRoutes(router)

	router.Run(addr)
}
