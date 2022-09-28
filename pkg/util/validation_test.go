package util_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/ghodss/yaml"
	v1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestValidation(t *testing.T) {
	t.Parallel()

	path := filepath.Join("test_data", "good_env.yaml")
	data, err := ioutil.ReadFile(path)
	require.NoError(t, err, "failed to load %s", path)

	deploy := &v1.Environment{}
	err = yaml.Unmarshal(data, deploy)
	require.NoError(t, err, "failed to unmarshal %s", path)

	results, err := util.ValidateYaml(deploy, data)
	t.Logf("got results %#v\n", results)

	require.NoError(t, err, "should not have failed to validate yaml file %s", path)

	require.Empty(t, results, "should not have validation errors for file %s", path)
}

func TestValidateSourceRepository(t *testing.T) {
	t.Parallel()

	path := filepath.Join("test_data", "good_sr.yaml")
	data, err := ioutil.ReadFile(path)
	require.NoError(t, err, "failed to load %s", path)

	deploy := &v1.SourceRepository{}
	err = yaml.Unmarshal(data, deploy)
	require.NoError(t, err, "failed to unmarshal %s", path)

	results, err := util.ValidateYaml(deploy, data)
	t.Logf("got results %#v\n", results)

	require.NoError(t, err, "should not have failed to validate yaml file %s", path)

	require.Empty(t, results, "should not have validation errors for file %s", path)
}

func TestValidateRelease(t *testing.T) {
	t.Parallel()

	path := filepath.Join("test_data", "good_release.yaml")
	data, err := ioutil.ReadFile(path)
	require.NoError(t, err, "failed to load %s", path)

	deploy := &v1.Release{}
	err = yaml.Unmarshal(data, deploy)
	require.NoError(t, err, "failed to unmarshal %s", path)

	results, err := util.ValidateYaml(deploy, data)
	t.Logf("got results %#v\n", results)

	require.NoError(t, err, "should not have failed to validate yaml file %s", path)

	require.Empty(t, results, "should not have validation errors for file %s", path)
}

func TestValidationFails(t *testing.T) {
	t.Parallel()

	path := filepath.Join("test_data", "bad_env.yaml")
	data, err := ioutil.ReadFile(path)
	require.NoError(t, err, "failed to load %s", path)

	deploy := &v1.Environment{}
	err = yaml.Unmarshal(data, deploy)
	require.NoError(t, err, "failed to unmarshal %s", path)

	results, err := util.ValidateYaml(deploy, data)
	t.Logf("got results %#v\n", results)

	require.NoError(t, err, "should not have failed to validate yaml file %s", path)

	require.NotEmpty(t, results, "should have validation errors for file %s", path)
}
