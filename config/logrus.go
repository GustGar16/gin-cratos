package config

import (
	//"os"
	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetFormatter(&logger.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	//TODO: checar si se debe almacenar en archivo o solo mostrar en consola
	//file, _ := os.OpenFile("apllication.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//logger.SetOutput(file)
}
func CreateLog(logType string, message string, function string) bool {
	log := logger.WithField("function", function)
	switch logType {
	case "info":
		log.Info(message)

	case "debug":
		logger.Debug(message)

	case "trace":
		logger.Trace(message)

	case "warn":
		logger.Warn(message)

	case "error":
		logger.Error(message)
	}
	return true
}
