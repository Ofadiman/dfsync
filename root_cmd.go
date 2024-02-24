package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

const (
	// Flag can take the following values:
	//	* `no-action` - Symlink will not be created for that particular file.
	//	* `backup` - Existing file will be moved to a file with `*.bak` extension and symlink will be created.
	//	* `override` - Existing file will be deleted and symlink will be created.
	OPTION_CONFLICT_RESOLUTION       = "conflict-resolution"
	OPTION_CONFLICT_RESOLUTION_SHORT = "c"
	OPTION_SOURCE_DIRECTORY          = "source-directory"
	OPTION_SOURCE_DIRECTORY_SHORT    = "s"
	OPTION_DRY                       = "dry"
	OPTION_DRY_SHORT                 = "d"
)

var validConflictResolutionFlagValues = []string{"no-action", "backup", "override"}

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
			conflictResolutionFlag := cmd.Flag(OPTION_CONFLICT_RESOLUTION)
			logger.Debugf("command has been called with the following flags: --%v=%v, --%v=%v, --%v=%v", OPTION_SOURCE_DIRECTORY, sourceFlag.Value.String(), OPTION_DRY, dryFlag.Value.String(), OPTION_CONFLICT_RESOLUTION, conflictResolutionFlag.Value.String())

			if slices.Contains[[]string, string](validConflictResolutionFlagValues, conflictResolutionFlag.Value.String()) == false {
				logger.Errorf("invalid value passed to --%v flag, received %v", OPTION_CONFLICT_RESOLUTION, conflictResolutionFlag.Value.String())
				return
			}

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

			absolute := getAbsolutePath(sourceFlag.Value.String())
			logger.Debugf("absolute path to source directory: %v", absolute)

			filepath.WalkDir(getAbsolutePath(sourceFlag.Value.String()), func(source string, d fs.DirEntry, err error) error {
				println()
				logger.Debugf("processing path: %v", source)
				if source == absolute {
					logger.Debugf("visiting source directory, no action required")
					return nil
				}

				trimmed := strings.TrimPrefix(source, getAbsolutePath(sourceFlag.Value.String())+"/")
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
								return nil
							}
						}
					}

					return nil
				}

				// TODO: Check how this command would behave if target path is symbolik link.
				_, err = os.Stat(target)
				if err != nil && errors.Is(err, os.ErrNotExist) == false {
					logger.Errorf("unhandled error: %v", err)
					return nil
				}

				if err == nil {
					if conflictResolutionFlag.Value.String() == "no-action" {
						logger.Warnf("path %v already exists, symlink will not be created because --%v flag is set to %v", target, OPTION_CONFLICT_RESOLUTION, conflictResolutionFlag.Value.String())
						return nil
					}

					if conflictResolutionFlag.Value.String() == "backup" {
						logger.Warnf("path %v already exists, the file will be backed up and symlink will be created because --%v flag is set to %v", target, OPTION_CONFLICT_RESOLUTION, conflictResolutionFlag.Value.String())

						backupPath := target + ".bak"

						if dryFlag.Value.String() == "false" {
							err := os.Rename(target, backupPath)
							if err != nil {
								logger.Errorf("unhandled error: %v", err)
								return nil
							}
						}

						logger.Infof("file has been successfully backup up under %v", backupPath)
					}

					if conflictResolutionFlag.Value.String() == "override" {
						logger.Warnf("path %v already exists, the file will be deleted and symlink will be created because --%v flag is set to %v", target, OPTION_CONFLICT_RESOLUTION, conflictResolutionFlag.Value.String())

						if dryFlag.Value.String() == "false" {
							err := os.Remove(target)
							if err != nil {
								logger.Errorf("unhandled error: %v", err)
								return nil
							}
						}
					}
				}

				cmd := exec.Command("ln", "-s", "-f", source, target)

				if dryFlag.Value.String() == "true" {
					logger.Infof("symlink from \"%v\" to \"%v\" created", source, target)
					return nil
				}

				output, err := cmd.CombinedOutput()
				if err != nil {
					logger.Debugf("am i here?")
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
	command.PersistentFlags().StringP(OPTION_CONFLICT_RESOLUTION, OPTION_CONFLICT_RESOLUTION_SHORT, "no-action", fmt.Sprintf("decide what to do with the file that exists where the symlink should be created (valid options: %v)", strings.Join(mapper[string, string](validConflictResolutionFlagValues, func(s string) string {
		return fmt.Sprintf("\"%v\"", s)
	}), ", ")))

	return command
}
