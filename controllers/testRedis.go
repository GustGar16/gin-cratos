package controllers

import (
	"fmt"
	"gin-Cratos/config"
)

func RedisTest() {
	config.SetRedisData("llave1", "prueba")
	val := config.GetRedisData("llave2")
	fmt.Println(val)
	config.CreateLog("info", "Llave en redis almacenada", "controller/testRedis::RedisTest()")
}
