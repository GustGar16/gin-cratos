package routes

import (
	"gin-Cratos/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine) {
	controllers.RedisTest()
	router.POST("/cargo", controllers.NewCargo())
	router.POST("/cancelacion/:uuid", controllers.NewCancelacion())

	router.GET("/transaccion/:uuid", controllers.GetTransaccion())
}
