package main

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

var logLevels = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"error": log.ErrorLevel,
	"warn":  log.WarnLevel,
	"fatal": log.FatalLevel,
}

func createLogger(out io.Writer) *log.Logger {
	logLevel := log.InfoLevel

	value, ok := logLevels[os.Getenv("LOG_LEVEL")]
	if ok {
		logLevel = value
	}

	return log.NewWithOptions(
		out, log.Options{
			Level:           logLevel,
			ReportTimestamp: false,
		},
	)
}
