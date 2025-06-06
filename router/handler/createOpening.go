package handler

import (
	"net/http"

	"github.com/WesleiDev/api-go-estudo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Create Opening
// @Description Crete a new Job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	logger.InfoF("Request received: %+v", request)

	opening := schemas.Opening {
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	err := db.Table("openings").Create(&opening).Error

	if(err != nil){
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx,"create-opening", opening)
	
}