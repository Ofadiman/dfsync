package main

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

func createLogger(out io.Writer) *log.Logger {
	return log.NewWithOptions(
		out, log.Options{
			ReportTimestamp: false,
		},
	)
}

func main() {
	logger := createLogger(os.Stdout)

	rootCommand := createRootCommand(logger)
	if err := rootCommand.Execute(); err != nil {
		logger.Fatal(err)
	}
}
