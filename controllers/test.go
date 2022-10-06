package controllers

import (
	"fmt"
	"gin-Cratos/models/fields"
	"gin-Cratos/models/header"
	"gin-Cratos/models/request"
	"gin-Cratos/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

var mti string

func SetMTI(value string) bool {
	mti = value
	return true
}

func GetMTI() string {
	return mti
}

func createHeaders() {
	//headers
	headerMti := header.HeaderData{Type: "n", Size: 4, Fixed: true, Mandatory: true, Usage: "MTI"}
	headerBitmap := header.HeaderData{Type: "n", Size: 8, Fixed: true, Mandatory: true, Usage: "Bit Map Primary"}
	headers := header.HeaderElement{MTI: headerMti, BITMAP: headerBitmap}
	fmt.Println(headers)

	//Fields data
	fields.NewFieldElement()
	fmt.Println(fields.FieldElement)
}

func cancelMessage(sMti string) {
	SetMTI(sMti)
}

func GetRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		createHeaders()
		var request request.TestRequest
		err := c.BindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.GeneralResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]string{"error": err.Error()}})
			return
		}
		cancelMessage("0200")
		c.JSON(http.StatusInternalServerError, responses.GeneralResponse{Status: http.StatusInternalServerError, Message: "error", Data: "aaa"})

	}
}
