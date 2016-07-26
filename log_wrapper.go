package log

import (
	"io"

	log "github.com/Sirupsen/logrus"
	echolog "github.com/labstack/gommon/log"
)

// utility wrapper struct that makes the link between logrus and echo logging interface
type LoggerWrapper struct {
	*log.Logger
}

func (LoggerWrapper) SetLevel(echolog.Lvl) {
	// Do nothing, do not let the framework change log level
}

func (logWrapper LoggerWrapper) Debugj(jsonArg echolog.JSON) {
	logWrapper.Debug(jsonArg)
}

func (logWrapper LoggerWrapper) Errorj(jsonArg echolog.JSON) {
	logWrapper.Error(jsonArg)
}

func (logWrapper LoggerWrapper) Fatalj(jsonArg echolog.JSON) {
	logWrapper.Fatal(jsonArg)
}

func (logWrapper LoggerWrapper) Infoj(jsonArg echolog.JSON) {
	logWrapper.Info(jsonArg)
}

func (logWrapper LoggerWrapper) Printj(jsonArg echolog.JSON) {
	logWrapper.Print(jsonArg)
}

func (logWrapper LoggerWrapper) Warnj(jsonArg echolog.JSON) {
	logWrapper.Warn(jsonArg)
}

func (logwrapper LoggerWrapper) SetOutput(io.Writer) {
	// do nothing do not let the framework change the writer
}
