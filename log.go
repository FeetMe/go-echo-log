package log

import (
	"os"

	"github.com/labstack/echo"

	log "github.com/Sirupsen/logrus"
)

func InitEchoServerLog(e *echo.Echo, prod bool, logFilePath string) error {
	if prod {
		f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		log.SetLevel(log.InfoLevel)
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(f)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	// forward echo logs to logrus
	loggerWrapper := LoggerWrapper{log.StandardLogger()}
	e.SetLogger(loggerWrapper)

	// also retrieve http request logs
	e.Use(getLogrusMiddlewareHandler())

	return nil
}
