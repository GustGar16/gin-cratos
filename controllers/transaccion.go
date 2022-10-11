package controllers

import (
	"gin-Cratos/responses"
	"gin-Cratos/services/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTransaccion() gin.HandlerFunc {
	return func(c *gin.Context) {
		transaccionId := c.Param("uuid")

		respuesta, err := database.GetRequest(transaccionId)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: respuesta})
	}
}
