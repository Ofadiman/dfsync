package utils

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

	assert.Equal(suite.T(), getAbsolutePath("~"), home)
	assert.Equal(suite.T(), getAbsolutePath("~/"), home)
	assert.Equal(suite.T(), getAbsolutePath("~/foo"), filepath.Join(home, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("~/foo/"), filepath.Join(home, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("~/foo/bar"), filepath.Join(home, "foo/bar"))
	assert.Equal(suite.T(), getAbsolutePath("~/foo/bar/"), filepath.Join(home, "foo", "bar"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleAbsolutePaths() {
	assert.Equal(suite.T(), getAbsolutePath("/"), "/")
	assert.Equal(suite.T(), getAbsolutePath("/foo"), "/foo")
	assert.Equal(suite.T(), getAbsolutePath("/foo/"), "/foo")
	assert.Equal(suite.T(), getAbsolutePath("/foo/bar"), "/foo/bar")
	assert.Equal(suite.T(), getAbsolutePath("/foo/bar/"), "/foo/bar")
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePaths() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), getAbsolutePath("./"), cwd)
	assert.Equal(suite.T(), getAbsolutePath("./foo"), filepath.Join(cwd, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("./foo/"), filepath.Join(cwd, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("./foo/bar"), filepath.Join(cwd, "foo", "bar"))
	assert.Equal(suite.T(), getAbsolutePath("./foo/bar/"), filepath.Join(cwd, "foo", "bar"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePathsStartingWithDot() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), getAbsolutePath(""), cwd)
	assert.Equal(suite.T(), getAbsolutePath("foo"), filepath.Join(cwd, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("foo/"), filepath.Join(cwd, "foo"))
	assert.Equal(suite.T(), getAbsolutePath("foo/bar"), filepath.Join(cwd, "foo", "bar"))
	assert.Equal(suite.T(), getAbsolutePath("foo/bar/"), filepath.Join(cwd, "foo", "bar"))
}

func TestGetAbsolutePathSuite(t *testing.T) {
	suite.Run(t, &getAbsolutePathSuite{})
}
