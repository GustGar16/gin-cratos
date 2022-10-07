package routes

import (
	"gin-Cratos/controllers"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine) {
	router.POST("/sale", controllers.SaleMessage())
}
