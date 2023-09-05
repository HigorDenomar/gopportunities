package router

import (
	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/handler"

	"github.com/higordenomar/gopportunities/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.Init()

	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.POST("/opportunity", handler.CreateOpportunityHandler)

		v1.GET("/opportunity", handler.ShowOpportunityHandler)

		v1.GET("/opportunities", handler.ListOpportunitiesHandler)

		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)

		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
	}

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
