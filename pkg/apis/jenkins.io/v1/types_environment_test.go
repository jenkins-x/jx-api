package v1_test

import (
	"encoding/json"
	"os"
	"path"
	"testing"

	v1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"

	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/stretchr/testify/assert"
)

var (
	testEnvironmentDataDir = path.Join("test_data", "environment")
)

func TestGitPublic(t *testing.T) {
	var gitPublicTests = []struct {
		jsonFile          string
		expectedGitPublic bool
	}{
		{"git_public_nil_git_private_true.json", false},
		{"git_public_nil_git_private_false.json", true},
		{"git_public_false_git_private_nil.json", false},
		{"git_public_true_git_private_nil.json", true},
	}

	for _, testCase := range gitPublicTests {
		t.Run(testCase.jsonFile, func(t *testing.T) {
			content, err := os.ReadFile(path.Join(testEnvironmentDataDir, testCase.jsonFile))
			assert.NoError(t, err)

			env := v1.Environment{}

			_ = log.CaptureOutput(func() {
				err = json.Unmarshal(content, &env)
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedGitPublic, env.Spec.TeamSettings.GitPublic, "unexpected value for default repository visibility")
			})
		})
	}
}

func Test_GitPublic_and_GitPrivate_specified_throws_error(t *testing.T) {
	content, err := os.ReadFile(path.Join(testEnvironmentDataDir, "git_public_true_git_private_true.json"))
	assert.NoError(t, err)

	env := v1.Environment{}
	err = json.Unmarshal(content, &env)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only GitPublic should be used")
}
