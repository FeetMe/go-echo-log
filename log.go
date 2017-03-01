package echo_log

import (
	"os"

	"github.com/labstack/echo"

	log "github.com/Sirupsen/logrus"
)

// configures the echo server to use logrus to log http request
// on "dev" and "prod" environment it also configures logrus to log into a file
func InitEchoServerLog(e *echo.Echo, env, projectName string) error {
	if env == "prod" || env == "production" || env == "dev" || env == "development" {
		project := projectName + "-" + env
		logFilePath := "/var/log/" + project + "/" + project + ".log"
		f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		if env == "prod" {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.DebugLevel)
		}
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(f)
	} else {
		log.SetLevel(log.DebugLevel)
	}

	// also retrieve http request logs
	e.Use(getLogrusMiddlewareHandler())

	return nil
}
