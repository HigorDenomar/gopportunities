package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "LIST Opportunity",
	})
}
