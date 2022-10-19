package controllers

import (
	"gin-Cratos/request"
	"gin-Cratos/responses"
	"gin-Cratos/services/database"
	"gin-Cratos/services/procesador"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Cargo() gin.HandlerFunc {
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
		//Armado de mensaje junto con la validacion
		message, err := procesador.Sale(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}
		if !database.InsertRequest(request, message) {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": "Error al almacenar la informacion: " + request.UUID}})
			return
		}

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: message})

	}
}

func Cancelacion() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.CancelacionRequest
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
		//TODO
		//Armado de mensaje junto con la validacion
		request.UUID = c.Param("uuid")
		message, err := procesador.Reversal(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: message})

	}
}

func Reembolso() {

}

func Captura() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request request.CapturaRequest
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
		//TODO
		//Armado de mensaje junto con la validacion
		request.UUID = c.Param("uuid")
		message, err := procesador.Capture(request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.GeneralResponse{Status: http.StatusOK, Message: "success", Data: message})

	}
}
