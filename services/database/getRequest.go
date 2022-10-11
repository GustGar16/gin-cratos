package database

import (
	"encoding/json"
	"fmt"
	"gin-Cratos/config"
	"gin-Cratos/request"
)

func GetRequest(transaccionId string) (request.CargoRequest, error) {

	q := "SELECT request FROM transaccion WHERE uuid = ?"
	row := config.DB.QueryRow(q, transaccionId)

	respuesta := ""
	data := request.CargoRequest{}
	err := row.Scan(&respuesta)

	if err != nil {
		fmt.Printf("SQL Err: %s", err)
		return data, err
	}

	json.Unmarshal([]byte(respuesta), &data)

	return data, nil
}
