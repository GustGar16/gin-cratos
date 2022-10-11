package main

import (
	"gin-Cratos/config"
	"gin-Cratos/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//inicializamos la configuracion por default para la declaracion de rutas
	router := gin.Default()

	config.Connect()
	defer config.Close()
	//Rutas declaradas
	routes.MainRoute(router)
	//Iniciamos el ruteo en el puerto declarado
	router.Run("localhost:6000")
}
