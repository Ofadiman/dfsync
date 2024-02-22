package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ofadiman/dfsync/utils"
	"github.com/spf13/cobra"
)

const (
	OPTION_SOURCE_DIRECTORY       = "source-directory"
	OPTION_SOURCE_DIRECTORY_SHORT = "s"
	OPTION_DRY                    = "dry"
	OPTION_DRY_SHORT              = "d"
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
			dryFlag := cmd.Flag(OPTION_DRY)

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

			home, _ := os.UserHomeDir()
			logger.Debugf("home directory: %v", home)

			absolute := utils.GetAbsolutePath(sourceFlag.Value.String())
			logger.Debugf("absolute path to source directory: %v", absolute)

			filepath.WalkDir(utils.GetAbsolutePath(sourceFlag.Value.String()), func(source string, d fs.DirEntry, err error) error {
				if source == absolute {
					logger.Debugf("visiting source directory, no action required")
					return nil
				}

				trimmed := strings.TrimPrefix(source, utils.GetAbsolutePath(sourceFlag.Value.String())+"/")
				target := filepath.Join(home, trimmed)

				if d.IsDir() {
					_, err := os.Stat(target)
					if err != nil {
						if os.IsNotExist(err) {
							logger.Warnf("directory \"%v\" does not exist, creating it", target)

							if dryFlag.Value.String() == "true" {
								return nil
							}

							err := os.Mkdir(target, 0700)
							if err != nil {
								logger.Errorf("unhandled error: %v", err)
							}
						}

						logger.Errorf("unhandled error: %v", err)
					}

					return nil
				}

				cmd := exec.Command("ln", "-s", "-f", source, target)

				if dryFlag.Value.String() == "true" {
					logger.Infof("symlink from \"%v\" to \"%v\" created", source, target)
					return nil
				}

				output, err := cmd.CombinedOutput()
				if err != nil {
					logger.Errorf("unhandled error: %v", strings.TrimSpace(string(output)))
				} else {
					logger.Infof("symlink from \"%v\" to \"%v\" created", source, target)
				}

				return nil
			})
		},
	}

	command.PersistentFlags().StringP(OPTION_SOURCE_DIRECTORY, OPTION_SOURCE_DIRECTORY_SHORT, cwd, "directory from which symlinks will be created")
	command.PersistentFlags().BoolP(OPTION_DRY, OPTION_DRY_SHORT, false, "simulate the execution of the command without modifying the file system")

	return command
}
