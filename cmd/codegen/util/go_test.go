package util_test

import (
	"fmt"

	"github.com/jenkins-x/jx-api/cmd/codegen/util"

	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	gopath = "GOPATH"
)

func Test_ensure_gopath_set(t *testing.T) {
	tmpGoDir, err := ioutil.TempDir("", "jx-codegen-tests")
	if err != nil {
		assert.Fail(t, "unable to create test directory")
	}
	defer os.RemoveAll(tmpGoDir)
	err = os.Setenv(gopath, tmpGoDir)
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	err = util.EnsureGoPath()
	assert.NoError(t, err, "GOPATH should be set")
}

func Test_ensure_gopath_unset(t *testing.T) {
	err := os.Setenv(gopath, "")
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	err = util.EnsureGoPath()
	assert.Error(t, err, "GOPATH should not be set")
	assert.Equal(t, "GOPATH needs to be set", err.Error())
}

func Test_ensure_gopath_does_not_exist(t *testing.T) {
	err := os.Setenv(gopath, "snafu")
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	err = util.EnsureGoPath()
	assert.Error(t, err, "GOPATH should not be set")
	assert.Equal(t, "the GOPATH directory snafu does not exist", err.Error())
}

func Test_get_gopath(t *testing.T) {
	tmpGoDir, err := ioutil.TempDir("", "jx-codegen-tests")
	if err != nil {
		assert.Fail(t, "unable to create test directory")
	}
	defer os.RemoveAll(tmpGoDir)
	err = os.Setenv(gopath, tmpGoDir)
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	goPath := util.GoPath()
	assert.Equal(t, tmpGoDir, goPath)
}

func Test_get_gopath_unset_env(t *testing.T) {
	err := os.Setenv(gopath, "")
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	goPath := util.GoPath()
	assert.Equal(t, "", goPath)
}

func Test_get_gopath_multiple_elements(t *testing.T) {
	tmpGoDir, err := ioutil.TempDir("", "jx-codegen-tests")
	if err != nil {
		assert.Fail(t, "unable to create test directory")
	}
	defer os.RemoveAll(tmpGoDir)
	err = os.Setenv(gopath, fmt.Sprintf("%s%sfoo", tmpGoDir, string(os.PathListSeparator)))
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	goPath := util.GoPath()
	assert.Equal(t, tmpGoDir, goPath)
}

func Test_get_gopath_src(t *testing.T) {
	tmpGoDir, err := ioutil.TempDir("", "jx-codegen-tests")
	if err != nil {
		assert.Fail(t, "unable to create test directory")
	}
	defer os.RemoveAll(tmpGoDir)
	err = os.Setenv(gopath, tmpGoDir)
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	goPathSrc := util.GoPathSrc(tmpGoDir)
	assert.Equal(t, filepath.Join(tmpGoDir, "src"), goPathSrc)
}

func Test_get_gopath_bin(t *testing.T) {
	tmpGoDir, err := ioutil.TempDir("", "jx-codegen-tests")
	if err != nil {
		assert.Fail(t, "unable to create test directory")
	}
	defer os.RemoveAll(tmpGoDir)
	err = os.Setenv(gopath, tmpGoDir)
	if err != nil {
		assert.Fail(t, "unable to set env variable")
	}

	goPathBin := util.GoPathBin(tmpGoDir)
	assert.Equal(t, filepath.Join(tmpGoDir, "bin"), goPathBin)
}
