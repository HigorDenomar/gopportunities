package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/higordenomar/gopportunities/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successful", operation),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int32  `json:"errorCode"`
}

type CreateOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type DeleteOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}

type ShowOpportunityResponse struct {
	Message string                      `json:"message"`
	Data    schemas.OpportunityResponse `json:"data"`
}
