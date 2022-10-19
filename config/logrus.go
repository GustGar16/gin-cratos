package config

import (
	"fmt"
	"io"
	"os"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	logger "github.com/sirupsen/logrus"
)

var logrusMessage = make(map[string]interface{})

func init() {
	logger.SetFormatter(&logger.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logger.SetReportCaller(true)
	setDefaultGraylogInfo()
	//comunicacion con graylog
	hook := graylog.NewGraylogHook(GetEnv("GRAYLOG_HOST", "localhost:12201"), logrusMessage)
	logger.AddHook(hook)
	//Impresión de log en archivo y en consola
	file, err := os.OpenFile(GetEnv("CRATOS_LOG_FILE", "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("error al abrir archivo de log: " + err.Error())
	}
	logger.SetOutput(io.MultiWriter(file, os.Stdout))
}

func CreateLog(logType string, message string, function string) bool {
	switch logType {
	case "info":
		logger.WithField("full_message", message).Info("Cratos::INFO")

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

func setDefaultGraylogInfo() {
	logrusMessage["version"] = GetEnv("CRATOS_VERSION")
	logrusMessage["host"] = GetEnv("CRATOS_HOST")
}
