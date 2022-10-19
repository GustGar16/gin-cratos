package database

import (
	"encoding/json"
	"fmt"
	"gin-Cratos/classes"
	"gin-Cratos/config"
)

func GetMessageRequest(transaccionId string) (classes.SaleMessage, error) {

	q := "SELECT messageRequest FROM transaccion WHERE uuid = ?"
	row := config.DB.QueryRow(q, transaccionId)

	respuesta := ""
	data := classes.SaleMessage{}
	err := row.Scan(&respuesta)

	if err != nil {
		fmt.Printf("SQL Err: %s", err)
		return data, err
	}

	json.Unmarshal([]byte(respuesta), &data)

	return data, nil
}
