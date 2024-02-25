package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RootCommandSuite struct {
	suite.Suite
	logs              *bytes.Buffer
	command           *cobra.Command
	fileSystemBuilder *FileSystemBuilder
	home              string
}

func (suite *RootCommandSuite) BeforeTest() {
	suite.T().Setenv("LOG_LEVEL", "debug")

	home, _ := os.UserHomeDir()
	suite.home = home

	suite.fileSystemBuilder = &FileSystemBuilder{}

	logsBuffer := new(bytes.Buffer)
	logger := createLogger(logsBuffer)

	suite.logs = logsBuffer
	suite.command = createRootCommand(logger)
	suite.command.SetOut(suite.logs)
	suite.command.SetErr(suite.logs)
}

func (suite *RootCommandSuite) TestShouldShowHelpMenuWhenCalledWithoutArguments() {
	suite.BeforeTest()

	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenConflictResolutionFlagIsInvalid() {
	suite.BeforeTest()

	suite.command.SetArgs([]string{"--conflict-resolution", "bar"})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToSourceDirectoryOptionDoesNotExist() {
	suite.BeforeTest()

	suite.command.SetArgs([]string{"--source-directory", "/foo/bar/buzz"})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToSourceDirectoryOptionIsNotDirectory() {
	suite.BeforeTest()

	filePath := "/tmp/foo.txt"
	cleanup := suite.fileSystemBuilder.File(filePath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", filePath})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenSourceDirectoryIsEmpty() {
	suite.BeforeTest()

	emptyDirPath := "/tmp/source/"
	cleanup := suite.fileSystemBuilder.Directory(emptyDirPath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", emptyDirPath})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldDoNothingWhenFileExistsAndNoActionIsPassedToConflictResolutionOption() {
	suite.BeforeTest()

	sourceDirectoryPath := "/tmp/source/"
	sourceFilePath := "/tmp/source/foo.txt"
	homeFilePath := filepath.Join(suite.home, "foo.txt")
	cleanup := suite.fileSystemBuilder.Directory(sourceDirectoryPath).File(sourceFilePath).File(homeFilePath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", sourceDirectoryPath})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
	homeFileStat, _ := os.Lstat(homeFilePath)
	assert.NotEqual(suite.T(), os.ModeSymlink, homeFileStat.Mode()&os.ModeSymlink)
}

func (suite *RootCommandSuite) TestShouldCreateBackupFileWhenFileExistsAndBackupIsPassedToConflictResolutionOption() {
	suite.BeforeTest()

	sourceDirectoryPath := "/tmp/source/"
	sourceFilePath := "/tmp/source/foo.txt"
	homeFilePath := filepath.Join(suite.home, "foo.txt")
	homeBackupFilePath := homeFilePath + ".bak"
	cleanup := suite.fileSystemBuilder.Directory(sourceDirectoryPath).File(sourceFilePath).File(homeFilePath).Build()
	defer cleanup(homeBackupFilePath)

	suite.command.SetArgs([]string{"--source-directory", sourceDirectoryPath, "--conflict-resolution", "backup"})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
	homeFileStat, _ := os.Lstat(homeFilePath)
	assert.Equal(suite.T(), os.ModeSymlink, homeFileStat.Mode()&os.ModeSymlink)
	_, err := os.Lstat(homeBackupFilePath)
	assert.Nil(suite.T(), err)
}

func (suite *RootCommandSuite) TestShouldOverrideFileWhenFileExistsAndOverrideIsPassedToConflictResolutionOption() {
	suite.BeforeTest()

	sourceDirectoryPath := "/tmp/source/"
	sourceFilePath := "/tmp/source/foo.txt"
	homeFilePath := filepath.Join(suite.home, "foo.txt")
	cleanup := suite.fileSystemBuilder.Directory(sourceDirectoryPath).File(sourceFilePath).File(homeFilePath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", sourceDirectoryPath, "--conflict-resolution", "override"})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
	homeFileStat, _ := os.Lstat(homeFilePath)
	assert.Equal(suite.T(), os.ModeSymlink, homeFileStat.Mode()&os.ModeSymlink)
}

func (suite *RootCommandSuite) TestShouldDoNothingWhenDryModeIsOn() {
	suite.BeforeTest()

	sourceDirectoryPath := "/tmp/source/"
	nestedDirectoryPath := "/tmp/source/nested/"
	fooFileSourcePath := "/tmp/source/foo.txt"
	barFileSourcePath := "/tmp/source/nested/bar.txt"
	fooFileTargetPath := filepath.Join(suite.home, "foo.txt")
	barFileTargetPath := filepath.Join(suite.home, "nested", "bar.txt")
	cleanup := suite.fileSystemBuilder.Directory(sourceDirectoryPath).Directory(nestedDirectoryPath).File(fooFileSourcePath).File(barFileSourcePath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", sourceDirectoryPath, "--dry", "true"})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
	_, err1 := os.Lstat(fooFileTargetPath)
	assert.True(suite.T(), errors.Is(err1, os.ErrNotExist))
	_, err2 := os.Lstat(barFileTargetPath)
	assert.True(suite.T(), errors.Is(err2, os.ErrNotExist))
}

func (suite *RootCommandSuite) TestShouldDoNothingWhenSourceDirectoryIsEmpty() {
	suite.BeforeTest()

	sourceDirectoryPath := "/tmp/source/"
	nestedDirectoryPath := "/tmp/source/nested/"
	nestedDirectoryTargetPath := filepath.Join(suite.home, "nested")
	cleanup := suite.fileSystemBuilder.Directory(sourceDirectoryPath).Directory(nestedDirectoryPath).Build()
	defer cleanup()

	suite.command.SetArgs([]string{"--source-directory", sourceDirectoryPath})
	suite.command.Execute()

	snaps.MatchSnapshot(suite.T(), suite.logs.String())
	_, err := os.Lstat(nestedDirectoryTargetPath)
	assert.True(suite.T(), errors.Is(err, os.ErrNotExist))
}

func TestRootCommandSuite(t *testing.T) {
	suite.Run(t, &RootCommandSuite{})
}
