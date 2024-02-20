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
		Example: fmt.Sprintf("dfsync --%v ./src --%v ~/", OPTION_SOURCE_DIRECTORY, OPTION_TARGET_DIRECTORY),
		Long:    `Dot files sync is a tool that allows you to painlessly synchronize dot files across multiple environments.`,
		Version: "1.0.0",
		Run: func(cmd *cobra.Command, args []string) {
			sourceFlag := cmd.Flag(OPTION_SOURCE_DIRECTORY)
			targetFlag := cmd.Flag(OPTION_TARGET_DIRECTORY)

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

			targetEntityInfo, err := os.Stat(targetFlag.Value.String())
			if err != nil {
				if os.IsNotExist(err) {
					logger.Errorf("directory passed to --%v option does not exist, received %v", OPTION_TARGET_DIRECTORY, targetFlag.Value.String())
					return
				}

				logger.Errorf("unhandled error: %v", err.Error())
				return
			}

			if targetEntityInfo.IsDir() == false {
				logger.Errorf("path passed to --%v option is not directory, received %v", OPTION_TARGET_DIRECTORY, targetFlag.Value.String())
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

			logger.Infof("Files will be symlinked from " + sourceFlag.Value.String() + " directory to " + targetFlag.Value.String() + " directory")

			filepath.WalkDir(utils.GetAbsolutePath(sourceFlag.Value.String()), func(sourcePath string, d fs.DirEntry, err error) error {
				cwd, _ := os.Getwd()
				targetPath := filepath.Join(utils.GetAbsolutePath(targetFlag.Value.String()), strings.Join(strings.Split(strings.TrimPrefix(sourcePath, cwd), "/")[2:], "/"))

				if d.IsDir() {
					err := os.MkdirAll(targetPath, 0700)
					if err != nil {
						logger.Errorf("could not create directory %v", targetPath)
						logger.Errorf("reason: %v", err.Error())
					}
				} else {
					logger.Infof("creating symlink for %v", sourcePath)

					cmd := exec.Command("ln", "-s", "-f", sourcePath, targetPath)
					if err := cmd.Run(); err != nil {
						logger.Errorf("could not create symlink at %v", targetPath)
						logger.Errorf("reason: %v", err.Error())
					} else {
						logger.Infof("%v (created)", targetPath)
					}
				}

				return nil
			})
		},
	}

	command.PersistentFlags().StringP(OPTION_SOURCE_DIRECTORY, "s", cwd, "directory from which symlinks will be created")
	command.PersistentFlags().StringP(OPTION_TARGET_DIRECTORY, "t", home, "directory to which symlinks will be created")

	return command
}
