package controllers

import (
	"gin-Cratos/responses"
	"net/http"

	"gin-Cratos/services/mainStruct"
	"gin-Cratos/services/messages"

	"github.com/gin-gonic/gin"
)

func NewCancelacion() gin.HandlerFunc {
	return func(c *gin.Context) {

		mainStruct.New()
		SetMTI("0420")
		message, err := messages.ReversalMessageContruct(mti, c.Param("uuid"))
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: message})

	}
}
