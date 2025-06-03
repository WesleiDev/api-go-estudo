package handler

import (
	"fmt"
	"net/http"

	"github.com/WesleiDev/api-go-estudo/schemas"
	"github.com/gin-gonic/gin"
)

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