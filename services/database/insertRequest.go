package database

import (
	"encoding/json"
	"fmt"
	"gin-Cratos/classes"
	"gin-Cratos/config"
	"gin-Cratos/request"
)

func InsertRequest(request request.CargoRequest, messageRequest classes.SaleMessage) bool {
	bin := request.Tarjeta.Pan[0:5]
	terminacion := request.Tarjeta.Pan[12:15]
	request.Tarjeta.Pan = bin + "******" + terminacion
	request.Tarjeta.Cvv = "***"

	requestJson, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(requestJson))

	messageJson, err := json.Marshal(messageRequest)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(messageJson))

	insert, err := config.DB.Query("INSERT INTO transaccion VALUES ( ?, ?, ? )", request.UUID, requestJson, messageJson)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	return true
}
