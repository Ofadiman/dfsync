package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	OPTION_SOURCE_DIRECTORY       = "source-directory"
	OPTION_SOURCE_DIRECTORY_SHORT = "s"
)

func createRootCommand(logger *log.Logger) *cobra.Command {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	command := &cobra.Command{
		Use:     "dfsync",
		Short:   "Synchronize dot files",
		Example: fmt.Sprintf("dfsync --%v ./src", OPTION_SOURCE_DIRECTORY),
		Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			sourceFlag := cmd.Flag(OPTION_SOURCE_DIRECTORY)

			sourceEntityInfo, err := os.Stat(sourceFlag.Value.String())
			if err != nil && os.IsNotExist(err) {
				if os.IsNotExist(err) {
					logger.Errorf("path passed to --%v option does not exist, received %v", OPTION_SOURCE_DIRECTORY, sourceFlag.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if sourceEntityInfo.IsDir() == false {
				logger.Errorf("path passed to --%v option is not directory, received %v", OPTION_SOURCE_DIRECTORY, sourceFlag.Value.String())
				return
			}

			dir, err := os.Open(sourceFlag.Value.String())
			defer dir.Close()
			if err != nil {
				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			_, err = dir.Readdirnames(1)
			if err == io.EOF {
				logger.Errorf("directory passed to --%v option is empty, received %v", OPTION_SOURCE_DIRECTORY, sourceFlag.Value.String())
				return
			}
		},
	}

	command.PersistentFlags().StringP(OPTION_SOURCE_DIRECTORY, OPTION_SOURCE_DIRECTORY_SHORT, cwd, "directory from which symlinks will be created")

	return command
}
