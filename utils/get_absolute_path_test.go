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

	assert.Equal(suite.T(), home, GetAbsolutePath("~"))
	assert.Equal(suite.T(), home, GetAbsolutePath("~/"))
	assert.Equal(suite.T(), filepath.Join(home, "foo"), GetAbsolutePath("~/foo"))
	assert.Equal(suite.T(), filepath.Join(home, "foo"), GetAbsolutePath("~/foo/"))
	assert.Equal(suite.T(), filepath.Join(home, "foo/bar"), GetAbsolutePath("~/foo/bar"))
	assert.Equal(suite.T(), filepath.Join(home, "foo", "bar"), GetAbsolutePath("~/foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleAbsolutePaths() {
	assert.Equal(suite.T(), "/", GetAbsolutePath("/"))
	assert.Equal(suite.T(), "/foo", GetAbsolutePath("/foo"))
	assert.Equal(suite.T(), "/foo", GetAbsolutePath("/foo/"))
	assert.Equal(suite.T(), "/foo/bar", GetAbsolutePath("/foo/bar"))
	assert.Equal(suite.T(), "/foo/bar", GetAbsolutePath("/foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePaths() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), cwd, GetAbsolutePath("./"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), GetAbsolutePath("./foo"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), GetAbsolutePath("./foo/"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), GetAbsolutePath("./foo/bar"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), GetAbsolutePath("./foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePathsStartingWithDot() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), cwd, GetAbsolutePath(""))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), GetAbsolutePath("foo"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo"), GetAbsolutePath("foo/"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), GetAbsolutePath("foo/bar"))
	assert.Equal(suite.T(), filepath.Join(cwd, "foo", "bar"), GetAbsolutePath("foo/bar/"))
}

func (suite *getAbsolutePathSuite) TestShouldHandleRelativePathsStartingWithTwoDots() {
	cwd, _ := os.Getwd()

	assert.Equal(suite.T(), filepath.Clean(filepath.Join(cwd, "../")), GetAbsolutePath(".."))
	assert.Equal(suite.T(), filepath.Clean(filepath.Join(cwd, "../")), GetAbsolutePath("../"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo"), GetAbsolutePath("../foo"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo"), GetAbsolutePath("../foo/"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo", "bar"), GetAbsolutePath("../foo/bar"))
	assert.Equal(suite.T(), filepath.Join(filepath.Clean(filepath.Join(cwd, "../")), "foo", "bar"), GetAbsolutePath("../foo/bar/"))
}

func TestGetAbsolutePathSuite(t *testing.T) {
	suite.Run(t, &getAbsolutePathSuite{})
}
