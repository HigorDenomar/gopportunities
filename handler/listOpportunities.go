package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/schemas"
)

func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
