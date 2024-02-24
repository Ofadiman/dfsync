package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type getAbsolutePathSuite struct {
	suite.Suite
}

func (suite *getAbsolutePathSuite) TestShouldHandleHomePaths() {
	home, _ := os.UserHomeDir()

	assert.Equal(suite.T(), filepath.Clean(home), getAbsolutePath("~"))
	assert.Equal(suite.T(), filepath.Clean(home), getAbsolutePath("~/"))
	assert.Equal(suite.T(), filepath.Join(home, "foo"), getAbsolutePath("~/foo"))
	assert.Equal(suite.T(), filepath.Join(home, "foo"), getAbsolutePath("~/foo/"))
	assert.Equal(suite.T(), filepath.Join(home, "foo/bar"), getAbsolutePath("~/foo/bar"))
	assert.Equal(suite.T(), filepath.Join(home, "foo", "bar"), getAbsolutePath("~/foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleAbsolutePaths() {
	assert.Equal(suite.T(), "/", getAbsolutePath("/"))
	assert.Equal(suite.T(), "/foo", getAbsolutePath("/foo"))
	assert.Equal(suite.T(), "/foo", getAbsolutePath("/foo/"))
	assert.Equal(suite.T(), "/foo/bar", getAbsolutePath("/foo/bar"))
	assert.Equal(suite.T(), "/foo/bar", getAbsolutePath("/foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePaths() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), cwd, getAbsolutePath("./"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), getAbsolutePath("./foo"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), getAbsolutePath("./foo/"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), getAbsolutePath("./foo/bar"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), getAbsolutePath("./foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePathsStartingWithDot() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), cwd, getAbsolutePath(""))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), getAbsolutePath("foo"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), getAbsolutePath("foo/"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), getAbsolutePath("foo/bar"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), getAbsolutePath("foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePathsStartingWithTwoDots() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), filepath.Clean(filepath.Join(cwd, "../")), getAbsolutePath(".."))
	assert.Equal(suite.T(), filepath.Clean(filepath.Join(cwd, "../")), getAbsolutePath("../"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo"), getAbsolutePath("../foo"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo"), getAbsolutePath("../foo/"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo", "bar"), getAbsolutePath("../foo/bar"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo", "bar"), getAbsolutePath("../foo/bar/"))
}

func TestGetAbsolutePathSuite(t *testing.T) {
	suite.Run(t, &getAbsolutePathSuite{})
}
