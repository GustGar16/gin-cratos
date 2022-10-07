package controllers

import (
	"gin-Cratos/request"
	"gin-Cratos/responses"
	"net/http"

	"gin-Cratos/services/mainStruct"

	"gin-Cratos/services/messages"

	"github.com/gin-gonic/gin"
)

var mti string

func SetMTI(value string) bool {
	mti = value
	return true
}

func SaleMessage() gin.HandlerFunc {
	return func(c *gin.Context) {
		mainStruct.New()
		var request request.SaleRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}

		SetMTI("0200")
		respuesta := messages.SaleMessage(mti, request)
		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: respuesta})

	}
}
