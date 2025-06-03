package handler

import (
	"fmt"
	"net/http"

	"github.com/WesleiDev/api-go-estudo/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// @Summary Delete Opening
// @Description Delete Job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == ""{
		sendError(ctx, http.StatusBadRequest, 
		errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if db.First(&opening, id).Error != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if db.Delete(&opening).Error != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error delete opening with id %s not found", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)

}