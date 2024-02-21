package main

import (
	"os"
)

func main() {
	logger := createLogger(os.Stdout)

	rootCommand := createRootCommand(logger)
	if err := rootCommand.Execute(); err != nil {
		logger.Fatal(err)
	}
}
