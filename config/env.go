package config

import (
	"log"

	"github.com/joho/godotenv"
)

var myEnv map[string]string = readEnv()

func readEnv() map[string]string {
	read, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return read
}
