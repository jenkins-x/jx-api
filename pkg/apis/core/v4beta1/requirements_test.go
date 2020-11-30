package v4beta1_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ghodss/yaml"
	"github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1"
	"github.com/jenkins-x/jx-api/v4/pkg/cloud"
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/stretchr/testify/assert"
)

var (
	testDataDir = path.Join("test_data")
)

const (
	KindGitHub = "github"
)

func TestRequirementsConfigMarshalExistingFile(t *testing.T) {
	t.Parallel()

	dir, err := ioutil.TempDir("", "test-requirements-config-")
	assert.NoError(t, err, "should create a temporary config dir")

	expectedClusterName := "my-cluster"
	expectedSecretStorage := v4beta1.SecretStorageTypeVault
	expectedDomain := "cheese.co.uk"
	expectedPipelineUserName := "someone"
	expectedPipelineUserEmail := "someone@acme.com"

	file := filepath.Join(dir, v4beta1.RequirementsConfigFileName)
	requirementResource := v4beta1.NewRequirementsConfig()
	requirements := &requirementResource.Spec
	requirements.SecretStorage = expectedSecretStorage
	requirements.Cluster.ClusterName = expectedClusterName
	requirements.Ingress.Domain = expectedDomain
	requirements.PipelineUser = &v4beta1.UserNameEmailConfig{
		Username: expectedPipelineUserName,
		Email:    expectedPipelineUserEmail,
	}

	err = requirementResource.SaveConfig(file)
	assert.NoError(t, err, "failed to save file %s", file)

	requirementResource, fileName, err := v4beta1.LoadRequirementsConfig(dir, v4beta1.DefaultFailOnValidationError)
	assert.NoError(t, err, "failed to load requirements file in dir %s", dir)
	assert.FileExists(t, fileName)
	requirements = &requirementResource.Spec
	assert.Equal(t, expectedClusterName, requirements.Cluster.ClusterName, "requirements.ClusterName")
	assert.Equal(t, expectedSecretStorage, requirements.SecretStorage, "requirements.SecretStorage")
	assert.Equal(t, expectedDomain, requirements.Ingress.Domain, "requirements.Domain")

	// lets check we can load the file from a sub dir
	subDir := filepath.Join(dir, "subdir")
	requirementResource, fileName, err = v4beta1.LoadRequirementsConfig(subDir, v4beta1.DefaultFailOnValidationError)
	assert.NoError(t, err, "failed to load requirements file in subDir: %s", subDir)
	assert.FileExists(t, fileName)
	requirements = &requirementResource.Spec
	t.Logf("generated requirements file %s\n", fileName)

	assert.Equal(t, expectedClusterName, requirements.Cluster.ClusterName, "requirements.ClusterName")
	assert.Equal(t, expectedSecretStorage, requirements.SecretStorage, "requirements.SecretStorage")
	assert.Equal(t, expectedDomain, requirements.Ingress.Domain, "requirements.Domain")

	require.NotNil(t, requirements.PipelineUser, "requirements.PipelineUser")
	assert.Equal(t, expectedPipelineUserName, requirements.PipelineUser.Username, "requirements.PipelineUser.Username")
	assert.Equal(t, expectedPipelineUserEmail, requirements.PipelineUser.Email, "requirements.PipelineUser.Email")

}

func Test_OverrideRequirementsFromEnvironment_does_not_initialise_nil_structs(t *testing.T) {
	requirementResource, fileName, err := v4beta1.LoadRequirementsConfig(testDataDir, v4beta1.DefaultFailOnValidationError)
	assert.NoError(t, err, "failed to load requirements file in dir %s", testDataDir)
	assert.FileExists(t, fileName)
	requirements := &requirementResource.Spec
	requirements.OverrideRequirementsFromEnvironment(func(in string) (string, error) {
		return "", nil
	})

	tempDir, err := ioutil.TempDir("", "test-requirements-config")
	assert.NoError(t, err, "should create a temporary config dir")
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	err = requirementResource.SaveConfig(filepath.Join(tempDir, v4beta1.RequirementsConfigFileName))
	assert.NoError(t, err, "failed to save requirements file in dir %s", tempDir)

	_, fileName, err = v4beta1.LoadRequirementsConfig(tempDir, v4beta1.DefaultFailOnValidationError)
	assert.NoError(t, err, "failed to load requirements file in dir %s", testDataDir)
	assert.FileExists(t, fileName)
}

func Test_OverrideRequirementsFromEnvironment_populate_requirements_from_environment_variables(t *testing.T) {
	var overrideTests = []struct {
		envKey               string
		envValue             string
		expectedRequirements v4beta1.RequirementsConfig
	}{
		// RequirementsConfig
		{v4beta1.RequirementSecretStorageType, "vault", v4beta1.RequirementsConfig{SecretStorage: "vault"}},
		{v4beta1.RequirementRepository, "bucketrepo", v4beta1.RequirementsConfig{Repository: "bucketrepo"}},
		{v4beta1.RequirementWebhook, "prow", v4beta1.RequirementsConfig{Webhook: "prow"}},

		// ClusterConfig
		{v4beta1.RequirementClusterName, "my-cluster", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{ClusterName: "my-cluster"}}},
		{v4beta1.RequirementProject, "my-project", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{ProjectID: "my-project"}}},
		{v4beta1.RequirementZone, "my-zone", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{Zone: "my-zone"}}},
		{v4beta1.RequirementChartRepository, "my-chart-museum", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{ChartRepository: "my-chart-museum"}}},
		{v4beta1.RequirementRegistry, "my-registry", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{Registry: "my-registry"}}},
		{v4beta1.RequirementEnvGitOwner, "john-doe", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{EnvironmentGitOwner: "john-doe"}}},
		{v4beta1.RequirementKanikoServiceAccountName, "kaniko-sa", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{KanikoSAName: "kaniko-sa"}}},
		{v4beta1.RequirementEnvGitPublic, "true", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{EnvironmentGitPublic: true}}},
		{v4beta1.RequirementEnvGitPublic, "false", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{EnvironmentGitPublic: false}}},
		{v4beta1.RequirementEnvGitPublic, "", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{EnvironmentGitPublic: false}}},
		{v4beta1.RequirementGitPublic, "true", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{GitPublic: true}}},
		{v4beta1.RequirementGitPublic, "false", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{GitPublic: false}}},
		{v4beta1.RequirementGitPublic, "", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{GitPublic: false}}},
		{v4beta1.RequirementExternalDNSServiceAccountName, "externaldns-sa", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{ExternalDNSSAName: "externaldns-sa"}}},

		// VaultConfig
		{v4beta1.RequirementVaultName, "my-vault", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{Name: "my-vault"}}},
		{v4beta1.RequirementVaultServiceAccountName, "my-vault-sa", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{ServiceAccount: "my-vault-sa"}}},
		{v4beta1.RequirementVaultKeyringName, "my-keyring", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{Keyring: "my-keyring"}}},
		{v4beta1.RequirementVaultKeyName, "my-key", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{Key: "my-key"}}},
		{v4beta1.RequirementVaultBucketName, "my-bucket", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{Bucket: "my-bucket"}}},
		{v4beta1.RequirementVaultRecreateBucket, "true", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{RecreateBucket: true}}},
		{v4beta1.RequirementVaultRecreateBucket, "false", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{RecreateBucket: false}}},
		{v4beta1.RequirementVaultRecreateBucket, "", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{RecreateBucket: false}}},
		{v4beta1.RequirementVaultDisableURLDiscovery, "true", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{DisableURLDiscovery: true}}},
		{v4beta1.RequirementVaultDisableURLDiscovery, "false", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{DisableURLDiscovery: false}}},
		{v4beta1.RequirementVaultDisableURLDiscovery, "", v4beta1.RequirementsConfig{Vault: v4beta1.VaultConfig{DisableURLDiscovery: false}}},

		// Storage
		{v4beta1.RequirementStorageBackupURL, "gs://my-backup",
			v4beta1.RequirementsConfig{Storage: []v4beta1.StorageConfig{
				{
					Name: "backup",
					URL:  "gs://my-backup",
				},
			}}},

		{v4beta1.RequirementStorageLogsURL, "gs://my-logs",
			v4beta1.RequirementsConfig{Storage: []v4beta1.StorageConfig{
				{
					Name: "logs",
					URL:  "gs://my-logs",
				},
			}}},
		{v4beta1.RequirementStorageReportsURL, "gs://my-reports",
			v4beta1.RequirementsConfig{Storage: []v4beta1.StorageConfig{
				{
					Name: "reports",
					URL:  "gs://my-reports",
				},
			}}},
		{v4beta1.RequirementStorageRepositoryURL, "gs://my-repo",
			v4beta1.RequirementsConfig{Storage: []v4beta1.StorageConfig{
				{
					Name: "repository",
					URL:  "gs://my-repo",
				},
			}}},

		// GKEConfig
		{v4beta1.RequirementGkeProjectNumber, "my-gke-project", v4beta1.RequirementsConfig{Cluster: v4beta1.ClusterConfig{GKEConfig: &v4beta1.GKEConfig{ProjectNumber: "my-gke-project"}}}},
	}

	for _, overrideTest := range overrideTests {
		origEnvValue, origValueSet := os.LookupEnv(overrideTest.envKey)
		err := os.Setenv(overrideTest.envKey, overrideTest.envValue)
		assert.NoError(t, err)
		resetEnvVariable := func() {
			var err error
			if origValueSet {
				err = os.Setenv(overrideTest.envKey, origEnvValue)
			} else {
				err = os.Unsetenv(overrideTest.envKey)
			}
			if err != nil {
				log.Logger().Warnf("error resetting environment after test: %v", err)
			}
		}

		t.Run(overrideTest.envKey, func(t *testing.T) {
			actualRequirements := v4beta1.RequirementsConfig{}
			actualRequirements.OverrideRequirementsFromEnvironment(func(projectId string) (string, error) {
				return "", nil
			})

			assert.Equal(t, overrideTest.expectedRequirements, actualRequirements)
		})

		resetEnvVariable()
	}
}

func TestRequirementsConfigMarshalExistingFileKanikoFalse(t *testing.T) {
	t.Parallel()

	dir, err := ioutil.TempDir("", "test-requirements-config-")
	assert.NoError(t, err, "should create a temporary config dir")

	file := filepath.Join(dir, v4beta1.RequirementsConfigFileName)
	requirementsResource := v4beta1.NewRequirementsConfig()
	err = requirementsResource.SaveConfig(file)
	assert.NoError(t, err, "failed to save file %s", file)

	_, fileName, err := v4beta1.LoadRequirementsConfig(dir, v4beta1.DefaultFailOnValidationError)
	assert.NoError(t, err, "failed to load requirements file in dir %s", dir)
	assert.FileExists(t, fileName)

}

func TestRequirementsConfigMarshalInEmptyDir(t *testing.T) {
	t.Parallel()

	dir, err := ioutil.TempDir("", "test-requirements-config-empty-")
	assert.NoError(t, err)

	requirements, fileName, err := v4beta1.LoadRequirementsConfig(dir, v4beta1.DefaultFailOnValidationError)
	assert.Error(t, err)
	assert.Empty(t, fileName)
	assert.Nil(t, requirements)
}

func TestRequirementsConfigIngressAutoDNS(t *testing.T) {
	t.Parallel()

	requirementsResource := v4beta1.NewRequirementsConfig()
	requirements := &requirementsResource.Spec
	requirements.Ingress.Domain = "1.2.3.4.nip.io"
	assert.Equal(t, true, requirements.Ingress.IsAutoDNSDomain(), "requirements.Ingress.IsAutoDNSDomain() for domain %s", requirements.Ingress.Domain)

	requirements.Ingress.Domain = "foo.bar"
	assert.Equal(t, false, requirements.Ingress.IsAutoDNSDomain(), "requirements.Ingress.IsAutoDNSDomain() for domain %s", requirements.Ingress.Domain)

	requirements.Ingress.Domain = ""
	assert.Equal(t, false, requirements.Ingress.IsAutoDNSDomain(), "requirements.Ingress.IsAutoDNSDomain() for domain %s", requirements.Ingress.Domain)
}

func Test_unmarshalling_requirements_config_with_build_pack_configuration_succeeds(t *testing.T) {
	t.Parallel()

	requirements := v4beta1.NewRequirementsConfig()

	content, err := ioutil.ReadFile(path.Join(testDataDir, "build_pack_library.yaml"))
	assert.NoError(t, err)

	err = yaml.Unmarshal(content, requirements)
	assert.NoError(t, err)
}

func Test_marshalling_empty_requirements_config_creates_no_build_pack_configuration(t *testing.T) {
	t.Parallel()

	requirements := v4beta1.NewRequirementsConfig()
	data, err := yaml.Marshal(requirements)
	assert.NoError(t, err)
	assert.NotContains(t, string(data), "buildPacks")

	err = yaml.Unmarshal(data, requirements)
	assert.NoError(t, err)
}

func Test_marshalling_vault_config(t *testing.T) {
	t.Parallel()

	requirementsResource := v4beta1.NewRequirementsConfig()
	requirements := &requirementsResource.Spec
	requirements.Vault = v4beta1.VaultConfig{
		Name:                   "myVault",
		URL:                    "http://myvault",
		ServiceAccount:         "vault-sa",
		Namespace:              "jx",
		KubernetesAuthPath:     "kubernetes",
		SecretEngineMountPoint: "secret",
	}
	data, err := yaml.Marshal(requirements)
	assert.NoError(t, err)

	assert.Contains(t, string(data), "name: myVault")
	assert.Contains(t, string(data), "url: http://myvault")
	assert.Contains(t, string(data), "serviceAccount: vault-sa")
	assert.Contains(t, string(data), "namespace: jx")
	assert.Contains(t, string(data), "kubernetesAuthPath: kubernetes")
	assert.Contains(t, string(data), "secretEngineMountPoint: secret")
}

func Test_env_repository_visibility(t *testing.T) {
	t.Parallel()

	var gitPublicTests = []struct {
		yamlFile          string
		expectedGitPublic bool
	}{
		{"git_public_nil_git_private_true.yaml", false},
		{"git_public_nil_git_private_false.yaml", true},
		{"git_public_false_git_private_nil.yaml", false},
		{"git_public_true_git_private_nil.yaml", true},
	}

	for _, testCase := range gitPublicTests {
		t.Run(testCase.yamlFile, func(t *testing.T) {
			content, err := ioutil.ReadFile(path.Join(testDataDir, testCase.yamlFile))
			assert.NoError(t, err)

			requirementsResource := v4beta1.NewRequirementsConfig()
			requirements := &requirementsResource.Spec
			_ = log.CaptureOutput(func() {
				err = yaml.Unmarshal(content, requirements)
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedGitPublic, requirements.Cluster.EnvironmentGitPublic, "unexpected value for repository visibility")
			})
		})
	}
}

func TestMergeSave(t *testing.T) {
	t.Parallel()
	type TestSpec struct {
		Name           string
		Original       *v4beta1.RequirementsConfig
		Changed        *v4beta1.RequirementsConfig
		ValidationFunc func(orig *v4beta1.RequirementsConfig, ch *v4beta1.RequirementsConfig)
	}

	testCases := []TestSpec{
		{
			Name: "Merge Cluster Config Test",
			Original: &v4beta1.RequirementsConfig{
				Cluster: v4beta1.ClusterConfig{
					EnvironmentGitOwner:  "owner",
					EnvironmentGitPublic: false,
					GitPublic:            false,
					Provider:             cloud.GKE,
					ProjectID:            "project-id",
					ClusterName:          "cluster-name",
					Region:               "region",
					GitKind:              KindGitHub,
					GitServer:            KindGitHub,
				},
			},
			Changed: &v4beta1.RequirementsConfig{
				Cluster: v4beta1.ClusterConfig{
					EnvironmentGitPublic: true,
					GitPublic:            true,
					Provider:             cloud.GKE,
				},
			},
			ValidationFunc: func(orig *v4beta1.RequirementsConfig, ch *v4beta1.RequirementsConfig) {
				assert.True(t, orig.Cluster.EnvironmentGitPublic == ch.Cluster.EnvironmentGitPublic &&
					orig.Cluster.GitPublic == ch.Cluster.GitPublic &&
					orig.Cluster.ProjectID != ch.Cluster.ProjectID, "ClusterConfig validation failed")
			},
		},
		{
			Name: "Merge EnvironmentConfig slices Test",
			Original: &v4beta1.RequirementsConfig{
				Environments: []v4beta1.EnvironmentConfig{
					{
						Key:        "dev",
						Repository: "repo",
					},
					{
						Key: "production",
						Ingress: &v4beta1.IngressConfig{
							Domain: "domain",
						},
					},
					{
						Key: "staging",
						Ingress: &v4beta1.IngressConfig{
							Domain: "domainStaging",
							TLS: &v4beta1.TLSConfig{
								Email: "email",
							},
						},
					},
				},
			},
			Changed: &v4beta1.RequirementsConfig{
				Environments: []v4beta1.EnvironmentConfig{
					{
						Key:   "dev",
						Owner: "owner",
					},
					{
						Key: "production",
						Ingress: &v4beta1.IngressConfig{
							CloudDNSSecretName: "secret",
						},
					},
					{
						Key: "staging",
						Ingress: &v4beta1.IngressConfig{
							Domain: "newDomain",
							TLS: &v4beta1.TLSConfig{
								Enabled: true,
							},
						},
					},
					{
						Key: "ns2",
					},
				},
			},
			ValidationFunc: func(orig *v4beta1.RequirementsConfig, ch *v4beta1.RequirementsConfig) {
				assert.True(t, len(orig.Environments) == len(ch.Environments), "the environment slices should be of the same len")
				// -- Dev Environment's asserts
				devOrig, devCh := orig.Environments[0], ch.Environments[0]
				assert.True(t, devOrig.Owner == devCh.Owner && devOrig.Repository != devCh.Repository,
					"the dev environment should've been merged correctly")
				// -- Production Environment's asserts
				prOrig, prCh := orig.Environments[1], ch.Environments[1]
				assert.True(t, prOrig.Ingress.Domain == "domain" &&
					prOrig.Ingress.CloudDNSSecretName == prCh.Ingress.CloudDNSSecretName,
					"the production environment should've been merged correctly")
				// -- Staging Environmnet's asserts
				stgOrig, stgCh := orig.Environments[2], ch.Environments[2]
				assert.True(t, stgOrig.Ingress.Domain == stgCh.Ingress.Domain &&
					stgOrig.Ingress.TLS.Email == "email" && stgOrig.Ingress.TLS.Enabled == stgCh.Ingress.TLS.Enabled,
					"the staging environment should've been merged correctly")
			},
		},
		{
			Name: "Merge StorageConfig test",
			Original: &v4beta1.RequirementsConfig{
				Storage: []v4beta1.StorageConfig{
					{
						Name: "logs",
						URL:  "value1",
					},
					{
						Name: "repository",
						URL:  "value3",
					},
				},
			},
			Changed: &v4beta1.RequirementsConfig{
				Storage: []v4beta1.StorageConfig{
					{
						Name: "reports",
						URL:  "value2",
					},
				},
			},
			ValidationFunc: func(orig *v4beta1.RequirementsConfig, ch *v4beta1.RequirementsConfig) {
				assert.True(t, orig.GetStorageURL("logs") == "value1" &&
					orig.GetStorageURL("repository") == "value3" &&
					orig.GetStorageURL("reports") == "value2",
					"The storage configuration should've been merged correctly")
			},
		},
	}
	f, err := ioutil.TempFile("", "")
	assert.NoError(t, err)
	defer func() {
		err := util.DeleteFile(f.Name())
		if err != nil {
			t.Logf("unable to clean up, %s", err)
		}
	}()

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			r := v4beta1.Requirements{
				Spec: *tc.Original,
			}
			rChanged := &v4beta1.Requirements{
				Spec: *tc.Changed,
			}
			err = r.MergeSave(rChanged, f.Name())
			assert.NoError(t, err, "the merge shouldn't fail for case %s", tc.Name)
			tc.ValidationFunc(&r.Spec, &rChanged.Spec)
		})
	}
}

func Test_EnvironmentGitPublic_and_EnvironmentGitPrivate_specified_together_return_error(t *testing.T) {
	content, err := ioutil.ReadFile(path.Join(testDataDir, "git_public_true_git_private_true.yaml"))
	assert.NoError(t, err)

	requirementsResource := v4beta1.NewRequirementsConfig()
	requirements := &requirementsResource.Spec
	err = yaml.Unmarshal(content, requirements)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "only EnvironmentGitPublic should be used")
}

func Test_LoadRequirementsConfig(t *testing.T) {
	t.Parallel()

	var gitPublicTests = []struct {
		requirementsPath   string
		createRequirements bool
	}{
		{"a", false},
		{"a/b", false},
		{"a/b/c", false},
		{"e", true},
		{"e/f", true},
		{"e/f/g", true},
	}

	for _, testCase := range gitPublicTests {
		t.Run(testCase.requirementsPath, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "jx-test-load-requirements-config")
			require.NoError(t, err, "failed to create tmp directory")
			defer func() {
				_ = os.RemoveAll(dir)
			}()

			testPath := filepath.Join(dir, testCase.requirementsPath)
			err = os.MkdirAll(testPath, 0700)
			require.NoError(t, err, "unable to create test path %s", testPath)

			var expectedRequirementsFile string
			if testCase.createRequirements {
				expectedRequirementsFile = filepath.Join(testPath, v4beta1.RequirementsConfigFileName)
				dummyRequirementsData := []byte("webhook: prow\n")
				err := ioutil.WriteFile(expectedRequirementsFile, dummyRequirementsData, 0600)
				require.NoError(t, err, "unable to write requirements file %s", expectedRequirementsFile)
			}

			requirementsResource, requirementsFile, err := v4beta1.LoadRequirementsConfig(testPath, v4beta1.DefaultFailOnValidationError)

			if testCase.createRequirements {
				requirements := &requirementsResource.Spec
				require.NoError(t, err)
				assert.Equal(t, expectedRequirementsFile, requirementsFile)
				assert.Equal(t, string(requirements.Webhook), "prow")
			} else {
				require.Error(t, err)
				assert.Equal(t, "", requirementsFile)
				assert.Nil(t, requirementsResource)
			}
		})
	}
}

func TestLoadRequirementsConfig_load_invalid_yaml(t *testing.T) {
	testDir := path.Join(testDataDir, "jx-requirements-syntax-error")

	absolute, err := filepath.Abs(testDir)
	assert.NoError(t, err, "could not find absolute path of dir %s", testDataDir)

	_, _, err = v4beta1.LoadRequirementsConfig(testDir, v4beta1.DefaultFailOnValidationError)
	requirementsConfigPath := path.Join(absolute, v4beta1.RequirementsConfigFileName)
	assert.EqualError(t, err, fmt.Sprintf("validation failures in YAML file %s:\nenvironments.0: Additional property namespace is not allowed", requirementsConfigPath))
}

func TestBackwardsCompatibleRequirementsFile(t *testing.T) {
	t.Parallel()
	oldRequirementsDir := path.Join(testDataDir, "backwards_compatible_requirements_file", "old")
	newRequirementsDir := path.Join(testDataDir, "backwards_compatible_requirements_file", "new")

	validateRequirements(t, oldRequirementsDir)
	validateRequirements(t, newRequirementsDir)
}

func validateRequirements(t *testing.T, oldRequirementsDir string) {
	requirementsResource, fileName, err := v4beta1.LoadRequirementsConfig(oldRequirementsDir, true)

	assert.NoError(t, err, "failed to load old style jx-requirements.yml")
	requirements := &requirementsResource.Spec
	assert.NotEmpty(t, fileName, "requirements filename should not be empty")
	assert.NotNil(t, requirements, "requirements should not be empty")
	assert.Equal(t, v4beta1.WebhookTypeLighthouse, requirements.Webhook, "failed to find requirement")
	assert.Equal(t, "gs://logs-foo", requirementsResource.Spec.GetStorageURL("logs"), "failed to find storage logs URL")
	assert.Equal(t, "1055835833001", requirementsResource.Spec.Cluster.GKEConfig.ProjectNumber, "failed to find project id")
}

func TestStorageURLHelpers(t *testing.T) {
	r := v4beta1.RequirementsConfig{}

	r.AddOrUpdateStorageURL("beer", "wine")
	assert.Equal(t, "wine", r.GetStorageURL("beer"))

	r.AddOrUpdateStorageURL("foo", "bar")
	assert.Equal(t, "bar", r.GetStorageURL("foo"))

	r.AddOrUpdateStorageURL("foo", "cheese")
	assert.Equal(t, "cheese", r.GetStorageURL("foo"))

	r.RemoveStorageURL("beer")
	assert.Equal(t, "", r.GetStorageURL("beer"))

	r.RemoveStorageURL("foo")
	assert.Equal(t, "", r.GetStorageURL("foo"))
}
func TestGetRequirementsConfigFromTeamSettings(t *testing.T) {

	content, err := ioutil.ReadFile(path.Join(testDataDir, "get_req_team_settings", "boot_requirements.yaml"))
	assert.NoError(t, err)

	settings := &v4beta1.TeamSettings{
		BootRequirements: string(content),
	}

	req, err := v4beta1.GetRequirementsConfigFromTeamSettings(settings)
	assert.NoError(t, err)
	assert.Equal(t, "http://bucketrepo/bucketrepo/charts/", req.Cluster.ChartRepository)
}
