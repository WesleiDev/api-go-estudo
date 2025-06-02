package router

import (
	"github.com/WesleiDev/api-go-estudo/router/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine){

	//Inicialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1");

	v1.GET("/opening", handler.ShowOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.GET("/openings", handler.ListOpeningsHandler)
}