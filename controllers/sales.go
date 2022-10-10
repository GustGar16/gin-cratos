package controllers

import (
	"gin-Cratos/request"
	"gin-Cratos/responses"
	"net/http"

	"gin-Cratos/services/mainStruct"

	"gin-Cratos/services/messages"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var mti string
var validate = validator.New()

func SetMTI(value string) bool {
	mti = value
	return true
}

func NewCargo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.CargoRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}
		//Se valida la estructura del cargo
		if validationErr := validate.Struct(&request); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": validationErr.Error()}})
			return
		}
		mainStruct.New()
		SetMTI("0200")
		message := messages.SaleMessageContruct(mti, request)

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: message})

	}
}
