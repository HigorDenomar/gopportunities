package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary List opportunities
// @Description List all job opportunities
// @Tags Opportunities
// @Accept json
// @Produce json
// @Success 200 {object} ListOpportunityResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunities [get]
func ListOpportunitiesHandler(ctx *gin.Context) {
	opportunities := []schemas.Opportunity{}

	if err := db.Find(&opportunities).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opportunities")
		return
	}

	sendSuccess(ctx, "list-opportunities", opportunities)
}
