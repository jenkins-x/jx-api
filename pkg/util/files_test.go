package util_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FileExists_for_non_existing_file_returns_false(t *testing.T) {
	exists, err := util.FileExists("/foo/bar")
	assert.NoError(t, err)
	assert.False(t, exists)
}

func Test_FileExists_for_existing_file_returns_true(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "Test_FileExists_for_existing_file_returns_true")
	require.NoError(t, err, "failed to create temporary directory")
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	data := []byte("hello\nworld\n")
	testFile := filepath.Join(tmpDir, "hello.txt")
	err = ioutil.WriteFile(testFile, data, 0600)
	require.NoError(t, err, "failed to create test file %s", testFile)

	exists, err := util.FileExists(testFile)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func Test_FileExists_for_existing_directory_returns_false(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "Test_FileExists_for_existing_file_returns_true")
	require.NoError(t, err, "failed to create temporary directory")
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	exists, err := util.FileExists(tmpDir)
	assert.NoError(t, err)
	assert.False(t, exists)
}
