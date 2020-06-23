package util_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/jenkins-x/jx-api/pkg/util"

	"github.com/stretchr/testify/assert"
)

func testCommand(t *testing.T, file string) {
	startPath, err := filepath.Abs("")
	if err != nil {
		panic(err)
	}
	tempfile, err := os.Create(filepath.Join(startPath, "/test_data/scripts", file))
	assert.NoError(t, err)
	tempfile.Close()
	defer os.Remove(tempfile.Name())

	cmd := util.Command{
		Name:    getFailIteratorScript(),
		Dir:     filepath.Join(startPath, "/test_data/scripts"),
		Args:    []string{file, "100"},
		Timeout: 3 * time.Second,
	}

	res, err := cmd.RunWithoutRetry()

	assert.Error(t, err, "Run should exit with failure")
	assert.Equal(t, "FAILURE!", res)
	assert.Equal(t, true, cmd.DidError())
	assert.Equal(t, true, cmd.DidFail())
	assert.Equal(t, 1, len(cmd.Errors))
	assert.Equal(t, 1, cmd.Attempts())
}
func TestRunWithoutRetry(t *testing.T) {
	t.Parallel()

	tmpFileName := "test_run_without_retry.txt"
	testCommand(t, tmpFileName)
}

func TestRunVerbose(t *testing.T) {
	t.Parallel()

	tmpFileName := "test_run_verbose.txt"
	testCommand(t, tmpFileName)
}

func TestRunQuiet(t *testing.T) {
	t.Parallel()

	tmpFileName := "test_run_quiet.txt"
	testCommand(t, tmpFileName)
}

func getFailIteratorScript() string {
	ex := "fail_iterator.sh"
	if runtime.GOOS == "windows" {
		ex = "fail_iterator.bat"
	}
	return ex
}
