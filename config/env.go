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

func GetEnv(values ...string) string {
	total := len(values)
	if total == 0 {
		return ""
	} else if total == 1 {
		return validateEnvKey(values[0])
	} else {
		mainKey := validateEnvKey(values[0])
		if mainKey != "" {
			return mainKey
		} else {
			return values[1]
		}
	}
	return ""
}

func validateEnvKey(key string) string {
	val, ok := myEnv[key]
	if !ok {
		return ""
	}
	return val
}
