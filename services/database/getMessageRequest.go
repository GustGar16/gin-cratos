package database

import (
	"encoding/json"
	"fmt"
	"gin-Cratos/config"
	"gin-Cratos/request"
)

func GetMessageRequest(transaccionId string) (request.SaleRequest, error) {

	q := "SELECT messageRequest FROM transaccion WHERE uuid = ?"
	row := config.DB.QueryRow(q, transaccionId)

	respuesta := ""
	data := request.SaleRequest{}
	err := row.Scan(&respuesta)

	if err != nil {
		fmt.Printf("SQL Err: %s", err)
		return data, err
	}

	json.Unmarshal([]byte(respuesta), &data)

	return data, nil
}
