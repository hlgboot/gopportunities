package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlgboot/gopportunities/schemas"
)

// @BasePath /api/v1

// Create Opening godoc
// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreatingOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreatingOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Link:     request.Link,
		Location: request.Location,
		Company:  request.Company,
		Salary:   request.Salary,
		Remote:   *request.Remote,
	}

	logger.Infof("request received: %+v", request)

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
