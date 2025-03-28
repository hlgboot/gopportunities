package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hlgboot/gopportunities/schemas"
)

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
