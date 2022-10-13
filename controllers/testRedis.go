package controllers

import (
	"fmt"
	"gin-Cratos/config"
)

func RedisTest() {
	config.SetRedisData("llave1", "prueba")
	val := config.GetRedisData("llave2")
	fmt.Println(val)
}
