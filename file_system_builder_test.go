package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateFileSystemAndRemoveIt(t *testing.T) {
	builder := FileSystemBuilder{}
	cleanup := builder.Directory("/tmp/test/one").Directory("/tmp/test/two").File("/tmp/test/foo.txt").File("/tmp/test/one/bar.txt").File("/tmp/test/two/buzz.txt").File("/tmp/linked.txt").Symlink("/tmp/linked.txt", "/tmp/test/linked.txt").Build()

	testOne, _ := os.Lstat("/tmp/test/one")
	assert.Equal(t, true, testOne.IsDir())

	testTwo, _ := os.Lstat("/tmp/test/two")
	assert.Equal(t, true, testTwo.IsDir())

	testFooTxt, _ := os.Lstat("/tmp/test/foo.txt")
	assert.Equal(t, false, testFooTxt.IsDir())

	testOneBarTxt, _ := os.Lstat("/tmp/test/one/bar.txt")
	assert.Equal(t, false, testOneBarTxt.IsDir())

	testTwoBuzzTxt, _ := os.Lstat("/tmp/test/two/buzz.txt")
	assert.Equal(t, false, testTwoBuzzTxt.IsDir())

	linkedTxt, _ := os.Lstat("/tmp/linked.txt")
	assert.Equal(t, false, linkedTxt.IsDir())

	testlinkedTxt, _ := os.Lstat("/tmp/test/linked.txt")
	assert.Equal(t, false, testlinkedTxt.IsDir())
	assert.NotEqual(t, 0, testlinkedTxt.Mode()&os.ModeSymlink)

	cleanup()

	_, err := os.Lstat("/tmp/test/one")
	assert.Error(t, err)
	assert.Equal(t, "lstat /tmp/test/one: no such file or directory", err.Error())

	_, err = os.Lstat("/tmp/test/two")
	assert.Error(t, err)

	_, err = os.Lstat("/tmp/test/foo.txt")
	assert.Error(t, err)

	_, err = os.Lstat("/tmp/test/one/bar.txt")
	assert.Error(t, err)

	_, err = os.Lstat("/tmp/test/two/buzz.txt")
	assert.Error(t, err)

	_, err = os.Lstat("/tmp/linked.txt")
	assert.Error(t, err)

	_, err = os.Lstat("/tmp/test/linked.txt")
	assert.Error(t, err)
}
