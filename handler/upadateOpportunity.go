package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Update opportunity
// @Description Update a job opportunity
// @Tags Opportunities
// @Accept json
// @Produce json
// @Param id query string true "Opportunity identification"
// @Param opportunity body UpdateOpportunityRequest true "Opportunity data to Update"
// @Success 200 {object} UpdateOpportunityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opportunity [put]
func UpdateOpportunityHandler(ctx *gin.Context) {
	request := UpdateOpportunityRequest{}
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opportunity := schemas.Opportunity{}

	if err := db.First(&opportunity, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opportunity with id %s not found", id))
		return
	}

	if request.Role != "" {
		opportunity.Role = request.Role
	}

	if request.Company != "" {
		opportunity.Company = request.Company
	}

	if request.Location != "" {
		opportunity.Location = request.Location
	}

	if request.Link != "" {
		opportunity.Link = request.Link
	}

	if request.Remote != nil {
		opportunity.Remote = *request.Remote
	}

	if request.Salary > 0 {
		opportunity.Salary = request.Salary
	}

	if err := db.Save(&opportunity).Error; err != nil {
		logger.Errorf("error updating opportunity: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opportunity")
		return
	}

	sendSuccess(ctx, "update-opportunity", opportunity)
}
