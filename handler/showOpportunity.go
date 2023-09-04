package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpportunityHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "SHOW Opportunity",
	})
}
