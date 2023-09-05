package router

import (
	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.Init()
	v1 := router.Group("/api/v1")

	{
		v1.POST("/opportunity", handler.CreateOpportunityHandler)

		v1.GET("/opportunity", handler.ShowOpportunityHandler)

		v1.GET("/opportunities", handler.ListOpportunitiesHandler)

		v1.PUT("/opportunity", handler.UpdateOpportunityHandler)

		v1.DELETE("/opportunity", handler.DeleteOpportunityHandler)
	}
}
