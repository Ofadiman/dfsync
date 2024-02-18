package main

import (
	"bytes"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RootCommandSuite struct {
	suite.Suite
	logger *log.Logger
	logs   *bytes.Buffer
}

func (suite *RootCommandSuite) SetupTest() {
	buff := new(bytes.Buffer)
	suite.logs = buff
	suite.logger = createLogger(buff)
}

func (suite *RootCommandSuite) TestSmoke() {
	assert.True(suite.T(), true, "Always passes")
}

func (suite *RootCommandSuite) TestOutputs() {
	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{})

	err := command.Execute()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "INFO Files will be synchronized from /home/szymon/projects/dfsync directory to /home/szymon directory\n", suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldReturnErrorWhenFromDirectoryNotFound() {
	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--from", "memes"})

	command.Execute()
	assert.Equal(suite.T(), "ERRO directory passed to --from option does not exist, received memes\n", suite.logs.String())
}

func (suite *RootCommandSuite) TestShouldReturnErrorWhenToDirectoryNotFound() {
	command := createRootCommand(suite.logger)
	command.SetOut(suite.logs)
	command.SetErr(suite.logs)
	command.SetArgs([]string{"--to", "memes"})

	command.Execute()
	assert.Equal(suite.T(), "ERRO directory passed to --to option does not exist, received memes\n", suite.logs.String())
}

func TestRootCommandSuite(t *testing.T) {
	suite.Run(t, &RootCommandSuite{})
}
