package main

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	OPTION_SOURCE_DIRECTORY = "source-directory"
	OPTION_TARGET_DIRECTORY = "target-directory"
)

func createRootCommand(logger *log.Logger) *cobra.Command {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}

	command := &cobra.Command{
		Use:     "dfsync",
		Short:   "Synchronize dot files",
		Example: "dfsync --source-directory ./src --target-directory ~/",
		Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			from := cmd.Flag(OPTION_SOURCE_DIRECTORY)
			to := cmd.Flag(OPTION_TARGET_DIRECTORY)

			fromFileInfo, err := os.Stat(from.Value.String())
			if err != nil && os.IsNotExist(err) {
				if os.IsNotExist(err) {
					logger.Errorf("path passed to --source-directory option does not exist, received %v", from.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if fromFileInfo.IsDir() == false {
				logger.Errorf("path passed to --source-directory option is not directory, received %v", from.Value.String())
				return
			}

			toFileInfo, err := os.Stat(to.Value.String())
			if err != nil {
				if os.IsNotExist(err) {
					logger.Errorf("directory passed to --target-directory option does not exist, received %v", to.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if toFileInfo.IsDir() == false {
				logger.Errorf("path passed to --target-directory option is not directory, received %v", to.Value.String())
				return
			}

			dir, err := os.Open(from.Value.String())
			defer dir.Close()
			if err != nil {
				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			_, err = dir.Readdirnames(1)
			if err == io.EOF {
				logger.Errorf("directory passed to --source-directory option is empty, received %v", from.Value.String())
				return
			}

			logger.Infof("Files will be synchronized from " + from.Value.String() + " directory to " + to.Value.String() + " directory")
		},
	}

	command.PersistentFlags().StringP(OPTION_SOURCE_DIRECTORY, "f", cwd, "directory from which symlinks will be created")
	command.PersistentFlags().StringP(OPTION_TARGET_DIRECTORY, "t", home, "directory to which symlinks will be created")

	return command
}
