package config_test

import (
	"path"
	"strings"
	"testing"

	"github.com/jenkins-x/jx-api/v3/pkg/config"

	"github.com/stretchr/testify/assert"
)

func TestJenkinsXAppsUnmarshalling(t *testing.T) {
	apps, err := config.LoadApplicationsConfig(path.Join("test_data"))
	assert.NoError(t, err)

	// assert marshalling of a jx-apps.yaml
	assert.Equal(t, 4, len(apps.Applications))
	assert.Equal(t, "cert-manager", apps.Applications[3].Namespace)
}

func TestBadPhase(t *testing.T) {
	_, err := config.LoadApplicationsConfig(path.Join("test_data", "jx-apps-phase-bad"))
	assert.Error(t, err)
	assert.True(t, strings.HasPrefix(err.Error(), "failed to validate YAML file"))
}

func TestGoodPhase(t *testing.T) {
	apps, err := config.LoadApplicationsConfig(path.Join("test_data", "jx-apps-phase-good"))
	assert.NoError(t, err)
	assert.Equal(t, "velero", apps.Applications[0].Name)
	assert.Equal(t, config.PhaseSystem, apps.Applications[0].Phase)
	assert.Equal(t, "external-dns", apps.Applications[1].Name)
	assert.Equal(t, config.PhaseApps, apps.Applications[1].Phase)
}
