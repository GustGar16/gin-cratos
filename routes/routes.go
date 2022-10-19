package routes

import (
	"gin-Cratos/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine) {
	router.POST("/cargo", controllers.Cargo())
	router.POST("/cancelacion/:uuid", controllers.Cancelacion())
	/*
		router.POST("/reembolso/:uuid", controllers.Reembolso())
		router.POST("/captura/:uuid", controllers.Captura())
	*/
	router.GET("/transaccion/:uuid", controllers.GetTransaccion())
}
