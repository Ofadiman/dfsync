package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RootCommandSuite struct {
	suite.Suite
	logger       *log.Logger
	logs         *bytes.Buffer
	file         *os.File
	emptyDirPath string
}

func (suite *RootCommandSuite) BeforeTest() {
	buff := new(bytes.Buffer)
	suite.logs = buff
	suite.logger = createLogger(buff)

	file, err := os.CreateTemp("/tmp", "")
	if err != nil {
		suite.logger.Fatal(err)
	}
	suite.file = file

	emptyDirPath, err := os.MkdirTemp("/tmp", "")
	if err != nil {
		suite.logger.Fatal(err)
	}
	suite.emptyDirPath = emptyDirPath
}

func (suite *RootCommandSuite) AfterTest() {
	if err := os.Remove(suite.file.Name()); err != nil {
		suite.logger.Fatal(err)
	}

	if err := os.Remove(suite.emptyDirPath); err != nil {
		suite.logger.Fatal(err)
	}
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToFromOptionDoesNotExist() {
	suite.BeforeTest()

	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--from", "memes"})

	command.Execute()

	assert.Equal(suite.T(), "ERRO path passed to --from option does not exist, received memes\n", suite.logs.String())

	suite.AfterTest()
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToFromOptionIsNotDirectory() {
	suite.BeforeTest()

	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--from", suite.file.Name()})

	command.Execute()

	assert.Equal(suite.T(), "ERRO path passed to --from option is not directory, received "+suite.file.Name()+"\n", suite.logs.String())

	suite.AfterTest()
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToToOptionDoesNotExist() {
	suite.BeforeTest()

	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--to", "memes"})

	command.Execute()

	assert.Equal(suite.T(), "ERRO directory passed to --to option does not exist, received memes\n", suite.logs.String())

	suite.AfterTest()
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenPathPassedToToOptionIsNotDirectory() {
	suite.BeforeTest()

	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--to", suite.file.Name()})

	command.Execute()

	assert.Equal(suite.T(), "ERRO path passed to --to option is not directory, received "+suite.file.Name()+"\n", suite.logs.String())

	suite.AfterTest()
}

func (suite *RootCommandSuite) TestShouldExitEarlyWhenFromDirectoryIsEmpty() {
	suite.BeforeTest()

	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--from", suite.emptyDirPath})

	command.Execute()

	assert.Equal(suite.T(), "ERRO directory passed to --from option is empty, received "+suite.emptyDirPath+"\n", suite.logs.String())

	suite.AfterTest()
}

func TestRootCommandSuite(t *testing.T) {
	suite.Run(t, &RootCommandSuite{})
}
