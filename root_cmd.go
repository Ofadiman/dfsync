package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	FLAG_FROM = "from"
	FLAG_TO   = "to"
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
		Example: "dfsync --source ./src --destination ~/",
		Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			from := cmd.Flag(FLAG_FROM)
			to := cmd.Flag(FLAG_TO)

			fromFileInfo, err := os.Stat(from.Value.String())
			if err != nil && os.IsNotExist(err) {
				if os.IsNotExist(err) {
					logger.Errorf("path passed to --from option does not exist, received %v", from.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if fromFileInfo.IsDir() == false {
				logger.Errorf("path passed to --from option is not directory, received %v", from.Value.String())
				return
			}

			toFileInfo, err := os.Stat(to.Value.String())
			if err != nil {
				if os.IsNotExist(err) {
					logger.Errorf("directory passed to --to option does not exist, received %v", to.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if toFileInfo.IsDir() == false {
				logger.Errorf("path passed to --to option is not directory, received %v", to.Value.String())
				return
			}

			logger.Infof("Files will be synchronized from " + from.Value.String() + " directory to " + to.Value.String() + " directory")
		},
	}

	command.PersistentFlags().StringP(FLAG_FROM, "f", cwd, "directory from which symlinks will be created")
	command.PersistentFlags().StringP(FLAG_TO, "t", home, "directory to which symlinks will be created")

	return command
}
