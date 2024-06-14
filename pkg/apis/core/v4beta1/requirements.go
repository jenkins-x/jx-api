package v4beta1

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	v1 "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jenkins-x/jx-api/v4/pkg/cloud"
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"

	"dario.cat/mergo"
	"github.com/ghodss/yaml"
	"github.com/vrischmann/envconfig"
)

var (
	// autoDNSSuffixes the DNS suffixes of any auto-DNS services
	autoDNSSuffixes = []string{
		".nip.io",
		".xip.io",
		".beesdns.com",
	}
)

const (
	RequirementsName = "Requirements"

	// DefaultFailOnValidationError by default fail if validation fails when reading jx-requirements
	DefaultFailOnValidationError = true

	constTrue  = "true"
	constFalse = "false"

	// Replaces the optional requirement and making jx hardcoded, if folks try changing the namespace in a jx-requirements.yml file it is highly likely to fail
	DefaultNamespace = "jx"

	backupName     = "backup"
	reportsName    = "reports"
	logsName       = "logs"
	repositoryName = "repository"
)

const (
	// RequirementsConfigFileName is the name of the requirements configuration file
	RequirementsConfigFileName = "jx-requirements.yml"
	// RequirementClusterName is the cluster name
	RequirementClusterName = "JX_REQUIREMENT_CLUSTER_NAME"
	// RequirementProject is the cloudprovider project
	RequirementProject = "JX_REQUIREMENT_PROJECT"
	// RequirementZone zone the cluster is in
	RequirementZone = "JX_REQUIREMENT_ZONE"
	// RequirementEnvGitOwner the default git owner for environment repositories if none is specified explicitly
	RequirementEnvGitOwner = "JX_REQUIREMENT_ENV_GIT_OWNER"
	// RequirementEnvGitPublic sets the visibility of the environment repositories as private (subscription required for GitHub Organisations)
	RequirementEnvGitPublic = "JX_REQUIREMENT_ENV_GIT_PUBLIC"
	// RequirementGitPublic sets the visibility of the application repositories as private (subscription required for GitHub Organisations)
	RequirementGitPublic = "JX_REQUIREMENT_GIT_PUBLIC"
	// RequirementExternalDNSServiceAccountName the service account name for external dns
	RequirementExternalDNSServiceAccountName = "JX_REQUIREMENT_EXTERNALDNS_SA_NAME"
	// RequirementVaultName the name for vault
	RequirementVaultName = "JX_REQUIREMENT_VAULT_NAME"
	// RequirementVaultServiceAccountName the service account name for vault
	RequirementVaultServiceAccountName = "JX_REQUIREMENT_VAULT_SA_NAME"
	// RequirementVaultKeyringName the keyring name for vault
	RequirementVaultKeyringName = "JX_REQUIREMENT_VAULT_KEYRING_NAME"
	// RequirementVaultKeyName the key name for vault
	RequirementVaultKeyName = "JX_REQUIREMENT_VAULT_KEY_NAME"
	// RequirementVaultBucketName the vault name for vault
	RequirementVaultBucketName = "JX_REQUIREMENT_VAULT_BUCKET_NAME"
	// RequirementVaultRecreateBucket recreate the bucket that vault uses
	RequirementVaultRecreateBucket = "JX_REQUIREMENT_VAULT_RECREATE_BUCKET"
	// RequirementVaultDisableURLDiscovery override the default lookup of the Vault URL, could be incluster service or external ingress
	RequirementVaultDisableURLDiscovery = "JX_REQUIREMENT_VAULT_DISABLE_URL_DISCOVERY"
	// RequirementSecretStorageType the secret storage type
	RequirementSecretStorageType = "JX_REQUIREMENT_SECRET_STORAGE_TYPE"
	// RequirementKanikoServiceAccountName the service account name for kaniko
	RequirementKanikoServiceAccountName = "JX_REQUIREMENT_KANIKO_SA_NAME"
	// RequirementIngressTLSProduction use the lets encrypt production server
	RequirementIngressTLSProduction = "JX_REQUIREMENT_INGRESS_TLS_PRODUCTION"
	// RequirementChartRepository the helm chart repository for jx
	RequirementChartRepository = "JX_REQUIREMENT_CHART_REPOSITORY"
	// RequirementRegistry the container registry for jx
	RequirementRegistry = "JX_REQUIREMENT_REGISTRY"
	// RequirementRepository the artifact repository for jx
	RequirementRepository = "JX_REQUIREMENT_REPOSITORY"
	// RequirementWebhook the webhook handler for jx
	RequirementWebhook = "JX_REQUIREMENT_WEBHOOK"
	// RequirementStorageBackupURL backup storage url
	RequirementStorageBackupURL = "JX_REQUIREMENT_STORAGE_BACKUP_URL"
	// RequirementStorageLogsURL logs storage url
	RequirementStorageLogsURL = "JX_REQUIREMENT_STORAGE_LOGS_URL"
	// RequirementStorageReportsURL report storage url
	RequirementStorageReportsURL = "JX_REQUIREMENT_STORAGE_REPORTS_URL"
	// RequirementStorageRepositoryURL repository storage url
	RequirementStorageRepositoryURL = "JX_REQUIREMENT_STORAGE_REPOSITORY_URL"
	// RequirementGkeProjectNumber is the gke project number
	RequirementGkeProjectNumber = "JX_REQUIREMENT_GKE_PROJECT_NUMBER"
	// RequirementDevEnvApprovers contains the optional list of users to populate the dev env's OWNERS with
	RequirementDevEnvApprovers = "JX_REQUIREMENT_DEV_ENV_APPROVERS"
)

// ChartRepositoryType is the type of chart repository used for helm
type ChartRepositoryType string

const (
	// ChartRepositoryTypeNone no kind so implies a chart repository you push tarballs to like chart museum / nexus
	ChartRepositoryTypeNone ChartRepositoryType = ""
	// ChartRepositoryTypeOCI specifies that we use OCI (container images) to store charts
	ChartRepositoryTypeOCI ChartRepositoryType = "oci"
	// ChartRepositoryTypePages specifies that we use github pages (a branch in git) to store helm charts
	ChartRepositoryTypePages ChartRepositoryType = "pages"
)

// ChartRepositoryTypeValues the string values for the secret storage
var ChartRepositoryTypeValues = []string{string(ChartRepositoryTypeOCI), string(ChartRepositoryTypePages)}

// SecretStorageType is the type of storage used for secrets
type SecretStorageType string

const (
	// SecretStorageTypeVault specifies that we use vault to store secrets
	SecretStorageTypeVault SecretStorageType = "vault"
	// SecretStorageTypeLocal specifies that we use the local file system in
	// `~/.jx/localSecrets` to store secrets
	SecretStorageTypeLocal SecretStorageType = "local"
)

// SecretStorageTypeValues the string values for the secret storage
var SecretStorageTypeValues = []string{string(SecretStorageTypeLocal), string(SecretStorageTypeVault)}

// WebhookType is the type of a webhook strategy
type WebhookType string

const (
	// WebhookTypeLighthouse specifies that we use lighthouse for webhooks
	// see: https://github.com/jenkins-x/lighthouse
	WebhookTypeLighthouse WebhookType = "lighthouse"
)

// WebhookTypeValues the string values for the webhook types
var WebhookTypeValues = []string{string(WebhookTypeLighthouse)}

// RepositoryType is the type of a repository we use to store artifacts (jars, tarballs, npm packages etc)
type RepositoryType string

const (
	// RepositoryTypeArtifactory if you wish to use Artifactory as the artifact repository
	RepositoryTypeArtifactory RepositoryType = "artifactory"
	// RepositoryTypeBucketRepo if you wish to use bucketrepo as the artifact repository. see https://github.com/jenkins-x/bucketrepo
	RepositoryTypeBucketRepo RepositoryType = "bucketrepo"
	// RepositoryTypeNone if you do not wish to install an artifact repository
	RepositoryTypeNone RepositoryType = "none"
	// RepositoryTypeNexus if you wish to use Sonatype Nexus as the artifact repository
	RepositoryTypeNexus RepositoryType = "nexus"
)

// RepositoryTypeValues the string values for the repository types
var RepositoryTypeValues = []string{string(RepositoryTypeNone), string(RepositoryTypeBucketRepo), string(RepositoryTypeNexus), string(RepositoryTypeArtifactory)}

// IngressType is the type of a ingress strategy
type IngressType string

const (
	// IngressTypeNone if we have yet to define a ingress type
	IngressTypeNone IngressType = ""
	// IngressTypeIngress uses the kubernetes extensions/v1 Ingress resources to create ingress
	IngressTypeIngress IngressType = "ingress"
	// IngressTypeIstio uses istio VirtualService resources to implement ingress instead of the extensions/v1 Ingress resources
	IngressTypeIstio IngressType = "istio"
	// IngressTypeHTTPRoute uses the Ingress V2 / HTTPRoute resources - see: https://kubernetes-sigs.github.io/service-apis/http-routing/
	IngressTypeHTTPRoute IngressType = "httproute"
)

// IngressTypeValues the string values for the ingress types
var IngressTypeValues = []string{"ingress", "istio", "httproute"}

// Requirements represents a collection installation requirements for Jenkins X
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Requirements struct {
	metav1.TypeMeta `json:",inline"`

	// Spec the definition of the secret mappings
	Spec RequirementsConfig `json:"spec"`
}

// EnvironmentConfig configures the organisation and repository name of the git repositories for environments
type EnvironmentConfig struct {
	// Key is the key of the environment configuration
	Key string `json:"key,omitempty"`
	// Owner is the git user or organisation for the repository
	Owner string `json:"owner,omitempty"`
	// Repository is the name of the repository within the owner
	Repository string `json:"repository,omitempty"`
	// GitServer is the URL of the git server
	GitServer string `json:"gitServer,omitempty"`
	// GitKind is the kind of git server (github, bitbucketserver etc)
	GitKind string `json:"gitKind,omitempty"`
	// GitURL optional git URL for the git repository for the environment. If its not specified its generated from the
	// git server, kind, owner and repository
	GitURL string `json:"gitUrl,omitempty"`
	// Ingress contains ingress specific requirements
	Ingress *IngressConfig `json:"ingress,omitempty"`
	// RemoteCluster specifies this environment runs on a remote cluster to the development cluster
	RemoteCluster bool `json:"remoteCluster,omitempty"`
	// PromotionStrategy what kind of promotion strategy to use
	PromotionStrategy v1.PromotionStrategyType `json:"promotionStrategy,omitempty"`
	// Should pull requests be labeled so that if there is an existing pull request for the application it can be found and updated
	ReusePullRequest bool `json:"reusePullRequest,omitempty"`
	// Namespace is the target namespace for deploying resources in this environment.  Will default to "jx-{{ .Key }}" if omitted
	Namespace string `json:"namespace,omitempty"`
}

// IngressConfig contains dns specific requirements
type IngressConfig struct {
	// APIVersion optional Ingress API version to use. Otherwise defaults to v1
	APIVersion string `json:"apiVersion,omitempty"`
	// DNS is enabled
	ExternalDNS bool `json:"externalDNS,omitempty"`
	// CloudDNSSecretName secret name which contains the service account for external-dns and cert-manager issuer to
	// access the Cloud DNS service to resolve a DNS challenge
	CloudDNSSecretName string `json:"cloud_dns_secret_name,omitempty"`
	// Domain to expose ingress endpoints
	Domain string `json:"domain"`
	// Kind the kind of ingress used (ingress v1, ingress v2, istio etc)
	Kind IngressType `json:"kind,omitempty"`
	// IgnoreLoadBalancer if the nginx-controller LoadBalancer service should not be used to detect and update the
	// domain if you are using a dynamic domain resolver like `.nip.io` rather than a real DNS configuration.
	// With this flag enabled the `Domain` value will be used and never re-created based on the current LoadBalancer IP address.
	IgnoreLoadBalancer bool `json:"ignoreLoadBalancer,omitempty"`
	// NamespaceSubDomain the sub domain expression to expose ingress. Defaults to ".jx."
	NamespaceSubDomain string `json:"namespaceSubDomain"`
	// TLS enable automated TLS using certmanager
	TLS *TLSConfig `json:"tls,omitempty"`
	// Annotations optional annotations added to ingresses
	Annotations map[string]string `json:"annotations,omitempty"`
}

// TLSConfig contains TLS specific requirements
type TLSConfig struct {
	// TLS enabled
	Enabled bool `json:"enabled"`
	// Email address to register with services like LetsEncrypt
	Email string `json:"email"`
	// Production false uses self-signed certificates from the LetsEncrypt staging server, true enables the production
	// server which incurs higher rate limiting https://letsencrypt.org/docs/rate-limits/
	Production bool `json:"production"`
	// SecretName the name of the secret which contains the TLS certificate
	SecretName string `json:"secretName,omitempty"`
}

// StorageConfig contains dns specific requirements
type StorageConfig struct {
	// Name of the bucket
	Name string `json:"name"`
	// URL the cloud storage bucket URL such as 'gs://mybucket' or 's3://foo' or `azblob://thingy'
	// see https://jenkins-x.io/architecture/storage/
	URL string `json:"url"`
}

type AzureDNSConfig struct {
	TenantID       string `json:"tenantId,omitempty"`
	SubscriptionID string `json:"subscriptionId,omitempty"`
	ResourceGroup  string `json:"resourceGroup,omitempty"`
}

type AzureSecretConfig struct {
	KeyVaultName string `json:"keyVaultName,omitempty"`
}

type AzureStorageConfig struct {
	StorageAccountName string `json:"storageAccountName,omitempty"`
}

type AzureClusterNodesConfig struct {
	ClientID string `json:"clientID"`
}

// AzureConfig contains Azure specific requirements
type AzureConfig struct {
	// RegistrySubscription the registry subscription for defaulting the container registry.
	// Not used if you specify a Registry explicitly
	RegistrySubscription     string                   `json:"registrySubscription,omitempty"`
	AzureDNSConfig           *AzureDNSConfig          `json:"dns,omitempty"`
	AzureSecretStorageConfig *AzureSecretConfig       `json:"secretStorage,omitempty"`
	AzureStorageConfig       *AzureStorageConfig      `json:"storage,omitempty"`
	AzureClusterNodesConfig  *AzureClusterNodesConfig `json:"clusterNodes,omitempty"`
}

// GKEConfig contains GKE specific requirements
type GKEConfig struct {
	// ProjectNumber the unique project number GKE assigns to a project (required for workload identity).
	ProjectNumber string `json:"projectNumber,omitempty" envconfig:"JX_REQUIREMENT_GKE_PROJECT_NUMBER"`
}

// DestinationConfig the common cluster settings that can be specified in settings or requirements
type DestinationConfig struct {
	// ChartRepository the repository URL to deploy charts to
	ChartRepository string `json:"chartRepository,omitempty" envconfig:"JX_REQUIREMENT_CHART_REPOSITORY"`
	// ChartKind the chart repository kind (e.g. normal, OCI or github pages)
	ChartKind ChartRepositoryType `json:"chartKind,omitempty" envconfig:"JX_REQUIREMENT_CHART_KIND"`
	// ChartSecret an optional secret name used to be able to push to chart repositories
	ChartSecret string `json:"chartSecret,omitempty" envconfig:"JX_REQUIREMENT_CHART_SECRET"`
	// Registry the host name of the container registry
	Registry string `json:"registry,omitempty" envconfig:"JX_REQUIREMENT_REGISTRY"`
	// DockerRegistryOrg the default organisation used for container images
	DockerRegistryOrg string `json:"dockerRegistryOrg,omitempty"`
	// KanikoFlags allows global kaniko flags to be supplied such as to disable host verification
	KanikoFlags string `json:"kanikoFlags,omitempty" envconfig:"JX_REQUIREMENT_KANIKO_FLAGS"`
	// EnvironmentGitOwner the default git owner for environment repositories if none is specified explicitly
	EnvironmentGitOwner string `json:"environmentGitOwner,omitempty" envconfig:"JX_REQUIREMENT_ENV_GIT_OWNER"`
}

// ClusterConfig contains cluster specific requirements
type ClusterConfig struct {
	DestinationConfig

	// AzureConfig the azure specific configuration
	AzureConfig *AzureConfig `json:"azure,omitempty"`
	// GKEConfig the gke specific configuration
	GKEConfig *GKEConfig `json:"gke,omitempty"`
	// EnvironmentGitPublic determines whether jx boot create public or private git repos for the environments
	EnvironmentGitPublic bool `json:"environmentGitPublic,omitempty" envconfig:"JX_REQUIREMENT_ENV_GIT_PUBLIC"`
	// GitPublic determines whether jx boot create public or private git repos for the applications
	GitPublic bool `json:"gitPublic,omitempty" envconfig:"JX_REQUIREMENT_GIT_PUBLIC"`
	// Provider the kubernetes provider (e.g. gke)
	Provider string `json:"provider,omitempty"`
	// ProjectID the cloud project ID e.g. on GCP
	ProjectID string `json:"project,omitempty" envconfig:"JX_REQUIREMENT_PROJECT"`
	// ClusterName the logical name of the cluster
	ClusterName string `json:"clusterName,omitempty" envconfig:"JX_REQUIREMENT_CLUSTER_NAME"`
	// Region the cloud region being used
	Region string `json:"region,omitempty"`
	// Zone the cloud zone being used
	Zone string `json:"zone,omitempty" envconfig:"JX_REQUIREMENT_ZONE"`
	// GitName is the name of the default git service
	GitName string `json:"gitName,omitempty"`
	// GitKind is the kind of git server (github, bitbucketserver etc)
	GitKind string `json:"gitKind,omitempty"`
	// GitServer is the URL of the git server
	GitServer string `json:"gitServer,omitempty"`
	// ExternalDNSSAName the service account name for external dns
	ExternalDNSSAName string `json:"externalDNSSAName,omitempty" envconfig:"JX_REQUIREMENT_EXTERNALDNS_SA_NAME"`
	// VaultSAName the service account name for vault
	// KanikoSAName the service account name for kaniko
	KanikoSAName string `json:"kanikoSAName,omitempty" envconfig:"JX_REQUIREMENT_KANIKO_SA_NAME"`
	// DevEnvApprovers contains an optional list of approvers to populate the initial OWNERS file in the dev env repo
	DevEnvApprovers []string `json:"devEnvApprovers,omitempty"`
	// Issue tracker to use for generating changelog
	IssueTracker *IssueTracker `json:"issueProvider,omitempty"`
}

// Deprecated: migrate to top level Requirements object
type legacyRequirementsConfig struct {
	RequirementsConfig `json:",inline"`

	Storage LegacyStorageConfig `json:"storage"`
}

// Deprecated: migrate to top level Requirements object
type LegacyStorageConfig struct {
	// Logs for storing build logs
	Logs LegacyStorageEntryConfig `json:"logs"`
	// Tests for storing test results, coverage + code quality reports
	Reports LegacyStorageEntryConfig `json:"reports"`
	// Repository for storing repository artifacts
	Repository LegacyStorageEntryConfig `json:"repository"`
	// Backup for backing up kubernetes resource
	Backup LegacyStorageEntryConfig `json:"backup"`
}

// Deprecated: migrate to top level Requirements object
type LegacyStorageEntryConfig struct {
	// Enabled if the storage is enabled
	Enabled bool `json:"enabled"`
	// URL the cloud storage bucket URL such as 'gs://mybucket' or 's3://foo' or `azblob://thingy'
	// see https://jenkins-x.io/architecture/storage/
	URL string `json:"url"`
}

// VaultConfig contains Vault configuration for Boot
type VaultConfig struct {
	// Name the name of the Vault if using Jenkins X managed Vault instance.
	// Cannot be used in conjunction with the URL attribute
	Name string `json:"name,omitempty"`

	Bucket         string `json:"bucket,omitempty" envconfig:"JX_REQUIREMENT_VAULT_BUCKET_NAME"`
	RecreateBucket bool   `json:"recreateBucket,omitempty"`

	Keyring string `json:"keyring,omitempty" envconfig:"JX_REQUIREMENT_VAULT_KEYRING_NAME"`
	Key     string `json:"key,omitempty" envconfig:"JX_REQUIREMENT_VAULT_KEY_NAME"`

	// DisableURLDiscovery allows us to optionally override the default lookup of the Vault URL, could be incluster service or external ingress
	DisableURLDiscovery bool `json:"disableURLDiscovery,omitempty"`

	// AWSConfig describes the AWS specific configuration needed for the Vault Operator.
	AWSConfig *VaultAWSConfig `json:"aws,omitempty"`

	// AzureConfig describes the Azure specific configuration needed for the Vault Operator.
	AzureConfig *VaultAzureConfig `json:"azure,omitempty"`

	// URL specifies the URL of an Vault instance to use for secret storage.
	// Needs to be specified together with the Service Account and namespace to use for connecting to Vault.
	// This cannot be used in conjunction with the Name attribute.
	URL string `json:"url,omitempty"`

	// ServiceAccount is the name of the Kubernetes service account allowed to authenticate against Vault.
	ServiceAccount string `json:"serviceAccount,omitempty" envconfig:"JX_REQUIREMENT_VAULT_SA_NAME"`

	// Namespace of the Kubernetes service account allowed to authenticate against Vault.
	Namespace string `json:"namespace,omitempty"`

	// SecretEngineMountPoint is the secret engine mount point to be used for writing data into the KV engine of Vault.
	// If not specified the 'secret' is used.
	SecretEngineMountPoint string `json:"secretEngineMountPoint,omitempty"`

	// KubernetesAuthPath is the auth path of used for this cluster
	// If not specified the 'kubernetes' is used.
	KubernetesAuthPath string `json:"kubernetesAuthPath,omitempty"`
}

// VaultAWSConfig contains all the Vault configuration needed by Vault to be deployed in AWS
type VaultAWSConfig struct {
	VaultAWSUnsealConfig
	AutoCreate          bool   `json:"autoCreate,omitempty"`
	DynamoDBTable       string `json:"dynamoDBTable,omitempty"`
	DynamoDBRegion      string `json:"dynamoDBRegion,omitempty"`
	ProvidedIAMUsername string `json:"iamUserName,omitempty"`
}

// VaultAWSUnsealConfig contains references to existing AWS resources that can be used to install Vault
type VaultAWSUnsealConfig struct {
	KMSKeyID  string `json:"kmsKeyId,omitempty"`
	KMSRegion string `json:"kmsRegion,omitempty"`
	S3Bucket  string `json:"s3Bucket,omitempty"`
	S3Prefix  string `json:"s3Prefix,omitempty"`
	S3Region  string `json:"s3Region,omitempty"`
}

// VaultAzureConfig contains all the Vault configuration needed by Vault to be deployed in Azure
type VaultAzureConfig struct {
	TenantID           string `json:"tenantId,omitempty"`
	VaultName          string `json:"vaultName,omitempty"`
	KeyName            string `json:"keyName,omitempty"`
	StorageAccountName string `json:"storageAccountName,omitempty"`
	ContainerName      string `json:"containerName,omitempty"`
}

// UnmarshalJSON method handles the rename of EnvironmentGitPrivate to EnvironmentGitPublic.
func (t *ClusterConfig) UnmarshalJSON(data []byte) error {
	// need a type alias to go into infinite loop
	type Alias ClusterConfig
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	_, gitPublicSet := raw["environmentGitPublic"]
	private, gitPrivateSet := raw["environmentGitPrivate"]

	if gitPrivateSet && gitPublicSet {
		return fmt.Errorf("found settings for EnvironmentGitPublic as well as EnvironmentGitPrivate in ClusterConfig, only EnvironmentGitPublic should be used")
	}

	if gitPrivateSet {
		log.Logger().Warn("EnvironmentGitPrivate specified in Cluster EnvironmentGitPrivate is deprecated use EnvironmentGitPublic instead.")
		privateString := string(private)
		if privateString == constTrue {
			t.EnvironmentGitPublic = false
		} else {
			t.EnvironmentGitPublic = true
		}
	}
	return nil
}

// AutoUpdateConfig contains auto update config
type AutoUpdateConfig struct {
	// Enabled autoupdate
	Enabled bool `json:"enabled"`
	// Schedule cron of auto updates
	Schedule string `json:"schedule"`
	// AutoMerge if enabled lets auto merge any generated update PullRequests on the dev cluster git repository
	AutoMerge bool `json:"autoMerge,omitempty"`
}

// RequirementsValues contains the logical installation requirements in the `jx-requirements.yml` file as helm values
type RequirementsValues struct {
	// RequirementsConfig contains the logical installation requirements
	RequirementsConfig *RequirementsConfig `json:"jxRequirements,omitempty"`
}

// UserNameEmailConfig contains the user name and email of a user (e.g. pipeline user)
type UserNameEmailConfig struct {
	// Username the username of the user
	Username string `json:"username,omitempty"`
	// Email the email address of the user
	Email string `json:"email,omitempty"`
}

// RequirementsConfig contains the logical installation requirements in the `jx-requirements.yml` file when
// installing, configuring or upgrading Jenkins X via `jx boot`
type RequirementsConfig struct {
	// AutoUpdate contains auto update config
	AutoUpdate AutoUpdateConfig `json:"autoUpdate,omitempty"`
	// Cluster contains cluster specific requirements
	Cluster ClusterConfig `json:"cluster"`
	// Environments the requirements for the environments
	Environments []EnvironmentConfig `json:"environments,omitempty"`
	// ExtraDomains to expose alternate services with custom ingress for specific applications
	ExtraDomains []IngressConfig `json:"extraDomains,omitempty"`
	// Ingress contains ingress specific requirements
	Ingress IngressConfig `json:"ingress"`
	// Kuberhealthy indicates if we have already installed Kuberhealthy upfront in the kubernetes cluster
	Kuberhealthy bool `json:"kuberhealthy,omitempty"`
	// PipelineUser the user name and email used for running pipelines
	PipelineUser *UserNameEmailConfig `json:"pipelineUser,omitempty"`
	// Repository specifies what kind of artifact repository you wish to use for storing artifacts (jars, tarballs, npm modules etc)
	Repository RepositoryType `json:"repository,omitempty" envconfig:"JX_REQUIREMENT_REPOSITORY"`
	// Repositories the configuration for language specific repositories
	Repositories *RepositoryConfig `json:"repositories,omitempty"`
	// SecretStorage how should we store secrets for the cluster
	SecretStorage SecretStorageType `json:"secretStorage,omitempty" envconfig:"JX_REQUIREMENT_SECRET_STORAGE_TYPE"`
	// Storage contains storage requirements
	Storage []StorageConfig `json:"storage,omitempty"`
	// Terraform specifies if  we are managing the kubernetes cluster and cloud resources with Terraform
	Terraform bool `json:"terraform,omitempty"`
	// TerraformVault indicates whether Vault has been installed upfront by Terraform
	TerraformVault bool `json:"terraformVault,omitempty"`
	// Vault the configuration for vault
	Vault VaultConfig `json:"vault,omitempty"`
	// Webhook specifies what engine we should use for webhooks
	Webhook WebhookType `json:"webhook,omitempty"`
}

// RepositoryConfig contains optional language specific repository configurations
type RepositoryConfig struct {
	// Maven the username of the user
	Maven *MavenRepositoryConfig `json:"maven,omitempty"`
}

// MavenRepositoryConfig contains optional configuration for maven repository configuration
type MavenRepositoryConfig struct {
	// ReleaseURL the release distribution URL
	ReleaseURL string `json:"releaseUrl,omitempty"`
	// SnapshotURL the snapshop distribution URL
	SnapshotURL string `json:"snapshotUrl,omitempty"`
}

// NewRequirementsConfig creates a default configuration file
func NewRequirementsConfig() *Requirements {
	return &Requirements{
		TypeMeta: metav1.TypeMeta{
			Kind:       RequirementsName,
			APIVersion: SchemeGroupVersion.String(),
		},
		Spec: RequirementsConfig{
			SecretStorage: SecretStorageTypeLocal,
			Webhook:       WebhookTypeLighthouse,
		},
	}
}

// LoadRequirementsConfig loads the project configuration if there is a project configuration file
// if there is not a file called `jx-requirements.yml` in the given dir we will scan up the parent
// directories looking for the requirements file as we often run 'jx' steps in sub directories.
func LoadRequirementsConfig(dir string, failOnValidationErrors bool) (*Requirements, string, error) {
	absolute, err := filepath.Abs(dir)
	if err != nil {
		return nil, "", fmt.Errorf("creating absolute path: %w", err)
	}
	if absolute != "" && absolute != "." && absolute != "/" {
		fileName := filepath.Join(absolute, RequirementsConfigFileName)
		exists, err := util.FileExists(fileName)
		if err != nil {
			return nil, "", err
		}
		if exists {
			requirements, err := LoadRequirementsConfigFile(fileName, failOnValidationErrors)
			return requirements, fileName, err
		}
	}
	return nil, "", fmt.Errorf("jx-requirements.yml file not found")
}

func IsNewRequirementsFile(s string) bool {
	if strings.Contains(s, "apiVersion:") && strings.Contains(s, "kind:") && strings.Contains(s, "spec:") {
		return true
	}
	return false
}

// LoadRequirementsConfigFile loads a specific project YAML configuration file
func LoadRequirementsConfigFileNoDefaults(fileName string, failOnValidationErrors bool) (*Requirements, error) {

	return loadRequirements(fileName, failOnValidationErrors)
}

// LoadRequirementsConfigFile loads a specific project YAML configuration file
func LoadRequirementsConfigFile(fileName string, failOnValidationErrors bool) (*Requirements, error) {
	requirements, err := loadRequirements(fileName, failOnValidationErrors)
	if requirements != nil {
		requirements.Spec.addDefaults()
	}
	return requirements, err
}

func loadRequirements(fileName string, failOnValidationErrors bool) (*Requirements, error) {
	requirements := &Requirements{
		TypeMeta: metav1.TypeMeta{
			Kind:       RequirementsName,
			APIVersion: SchemeGroupVersion.String(),
		}}
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("checking if file %s exists: %w", fileName, err)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to load file %s due to %s", fileName, err)
	}

	// //check whether new or old jx requirements
	if IsNewRequirementsFile(string(data)) {
		validationErrors, err := util.ValidateYaml(requirements, data)
		if err != nil {
			return nil, fmt.Errorf("failed to validate YAML file %s due to %s", fileName, err)
		}

		if len(validationErrors) > 0 {
			log.Logger().Warnf("validation failures in YAML file %s: %s", fileName, strings.Join(validationErrors, ", "))
			if failOnValidationErrors {
				return nil, fmt.Errorf("validation failures in YAML file %s:\n%s", fileName, strings.Join(validationErrors, "\n"))
			}
		}

		err = yaml.Unmarshal(data, requirements)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML file %s due to %s", fileName, err)
		}

	} else {
		config := &legacyRequirementsConfig{}
		validationErrors, err := util.ValidateYaml(config, data)
		if err != nil {
			return nil, fmt.Errorf("failed to validate YAML file %s due to %s", fileName, err)
		}

		if len(validationErrors) > 0 {
			log.Logger().Warnf("validation failures in YAML file %s: %s", fileName, strings.Join(validationErrors, ", "))

			if failOnValidationErrors {
				return nil, fmt.Errorf("validation failures in YAML file %s:\n%s", fileName, strings.Join(validationErrors, "\n"))
			}
		}

		err = yaml.Unmarshal(data, config)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal YAML file %s due to %s", fileName, err)
		}

		requirements.migrateV3(config)
	}

	return requirements, nil
}

// IsEmpty returns true if this configuration is empty
func (c *RequirementsConfig) IsEmpty() bool {
	empty := &RequirementsConfig{}
	return reflect.DeepEqual(empty, c)
}

func (c *Requirements) SaveConfig(fileName string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, data, util.DefaultWritePermissions)
	if err != nil {
		return fmt.Errorf("failed to save file %s: %w", fileName, err)
	}

	return nil
}

type sliceTransformer struct{}

// sliceTransformer.Transformer is handling the correct merge of two EnvironmentConfig slices
// so we can both append extra items and merge existing ones so we don't lose any data
func (t sliceTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf([]EnvironmentConfig{}) {
		return func(dst, src reflect.Value) error {
			d := dst.Interface().([]EnvironmentConfig)
			s := src.Interface().([]EnvironmentConfig)
			if dst.CanSet() {
				for i := range s {
					v := s[i]
					if i > len(d)-1 {
						d = append(d, v)
					} else {
						nv := v
						err := mergo.Merge(&d[i], &nv, mergo.WithOverride)
						if err != nil {
							return fmt.Errorf("error merging EnvironmentConfig slices: %w", err)
						}
					}
				}
				dst.Set(reflect.ValueOf(d))
			}
			return nil
		}
	}
	if typ == reflect.TypeOf([]StorageConfig{}) {
		return func(dst, src reflect.Value) error {
			d := dst.Interface().([]StorageConfig)
			s := src.Interface().([]StorageConfig)
			if dst.CanSet() {
				for i, v := range s {
					if len(d) > 0 {
						d = append(d, v)
					} else {
						nv := v
						err := mergo.Merge(&d[i], &nv, mergo.WithOverride)
						if err != nil {
							return fmt.Errorf("error merging StorageConfig slices: %w", err)
						}
					}
				}
				dst.Set(reflect.ValueOf(d))
			}
			return nil
		}
	}
	return nil
}

// MergeSave attempts to merge the provided Requirements with the caller's data.
// It does so overriding values in the source struct with non-zero values from the provided struct
// it defines non-zero per property and not for a while embedded struct, meaning that nested properties
// in embedded structs should also be merged correctly.
// if a slice is added a transformer will be needed to handle correctly merging the contained values
func (c *Requirements) MergeSave(src *Requirements, requirementsFileName string) error {
	err := mergo.Merge(c, src, mergo.WithOverride, mergo.WithTransformers(sliceTransformer{}))
	if err != nil {
		return fmt.Errorf("error merging jx-requirements.yml files: %w", err)
	}
	err = c.SaveConfig(requirementsFileName)
	if err != nil {
		return fmt.Errorf("error saving the merged jx-requirements.yml files to %s: %w", requirementsFileName, err)
	}
	return nil
}

// lets remove this on the next major version update, currently v4
func (c *Requirements) migrateV3(config *legacyRequirementsConfig) {
	c.Spec = config.RequirementsConfig

	if config.Storage.Backup.Enabled && config.Storage.Backup.URL != "" {
		c.Spec.AddOrUpdateStorageURL(backupName, config.Storage.Backup.URL)
	}
	if config.Storage.Logs.Enabled && config.Storage.Logs.URL != "" {
		c.Spec.AddOrUpdateStorageURL(logsName, config.Storage.Logs.URL)
	}
	if config.Storage.Reports.Enabled && config.Storage.Reports.URL != "" {
		c.Spec.AddOrUpdateStorageURL(reportsName, config.Storage.Reports.URL)
	}
	if config.Storage.Repository.Enabled && config.Storage.Repository.URL != "" {
		c.Spec.AddOrUpdateStorageURL(repositoryName, config.Storage.Repository.URL)
	}
}

// EnvironmentMap creates a map of maps tree which can be used inside Go templates to access the environment
// configurations
func (c *RequirementsConfig) EnvironmentMap() map[string]interface{} {
	answer := map[string]interface{}{}
	for i := range c.Environments {
		env := c.Environments[i]
		k := env.Key
		if k == "" {
			log.Logger().Warnf("missing 'key' for Environment requirements %#v", env)
			continue
		}
		e := env
		m, err := util.ToObjectMap(&e)
		if err == nil {
			ensureHasFields(m, "owner", "repository", "gitServer", "gitKind")
			answer[k] = m
		} else {
			log.Logger().Warnf("failed to turn environment %s with value %#v into a map: %s\n", k, e, err.Error())
		}
	}
	return answer
}

// Environment looks up the environment configuration based on environment name
func (c *RequirementsConfig) Environment(name string) (*EnvironmentConfig, error) {
	for i := range c.Environments {
		env := c.Environments[i]
		if env.Key == name {
			return &env, nil
		}
	}
	return nil, fmt.Errorf("environment %q not found", name)
}

// ToMap converts this object to a map of maps for use in helm templating
func (c *RequirementsConfig) ToMap() (map[string]interface{}, error) {
	m, err := util.ToObjectMap(c)
	if m != nil {
		ensureHasFields(m, "provider", "project", "environmentGitOwner", "gitops", "webhook")
	}
	if m["repositories"] == nil {
		m["repositories"] = map[string]interface{}{
			"maven": map[string]interface{}{
				"releaseUrl":  "",
				"snapshotUrl": "",
			},
		}
	}
	return m, err
}

// IsCloudProvider returns true if the kubenretes provider is a cloud
func (c *RequirementsConfig) IsCloudProvider() bool {
	p := c.Cluster.Provider
	return p == cloud.GKE || p == cloud.AKS || p == cloud.AWS || p == cloud.EKS || p == cloud.ALIBABA
}

func ensureHasFields(m map[string]interface{}, keys ...string) {
	for _, k := range keys {
		_, ok := m[k]
		if !ok {
			m[k] = ""
		}
	}
}

// MissingRequirement returns an error if there is a missing property in the requirements
func MissingRequirement(property, fileName string) error {
	return fmt.Errorf("missing property: %s in file %s", property, fileName)
}

// IsLazyCreateSecrets returns a boolean whether secrets should be lazily created
func (c *RequirementsConfig) IsLazyCreateSecrets(flag string) (bool, error) {
	if flag != "" {
		switch flag {
		case constTrue:
			return true, nil
		case constFalse:
			return false, nil
		default:
			return false, fmt.Errorf("invalid option for lazy-create: %s", flag)
		}
	} else if !c.Terraform {
		return true, nil
	}

	// default to false
	return false, nil
}

// addDefaults lets ensure any missing values have good defaults
func (c *RequirementsConfig) addDefaults() {
	if c.Cluster.GitServer == "" {
		c.Cluster.GitServer = "https://github.com"
	}
	if c.Cluster.GitKind == "" {
		c.Cluster.GitKind = "github"
	}
	if c.Cluster.GitName == "" {
		c.Cluster.GitName = c.Cluster.GitKind
	}
	if c.Ingress.NamespaceSubDomain == "" {
		c.Ingress.NamespaceSubDomain = "-" + "jx" + "."
	}
	if c.Ingress.Kind == "" {
		c.Ingress.Kind = IngressTypeIngress
	}
	if c.Webhook == "" {
		c.Webhook = WebhookTypeLighthouse
	}
	if c.Repository == "" {
		c.Repository = "nexus"
	}
}

// IsAutoDNSDomain returns true if the domain is configured to use an auto DNS sub domain like
// '.nip.io' or '.xip.io'
func (i *IngressConfig) IsAutoDNSDomain() bool {
	for _, suffix := range autoDNSSuffixes {
		if strings.HasSuffix(i.Domain, suffix) {
			return true
		}
	}
	return false
}

// OverrideRequirementsFromEnvironment allows properties to be overridden with environment variables
func (c *RequirementsConfig) OverrideRequirementsFromEnvironment(gkeProjectNumber func(projectId string) (string, error)) {
	// struct members need to use explicit 'envconfig:"<var-name>"' unless there is a match between struct member navigation
	// path and env variable name
	err := envconfig.InitWithOptions(&c, envconfig.Options{AllOptional: true, LeaveNil: true, Prefix: "JX_REQUIREMENT"})
	if err != nil {
		log.Logger().Errorf("Unable to init envconfig for override requirements: %s", err)
	}

	// RequirementIngressTLSProduction applies to more than one environment and needs to be handled explicitly
	if os.Getenv(RequirementIngressTLSProduction) != "" {
		useProduction := os.Getenv(RequirementIngressTLSProduction)
		if envVarBoolean(useProduction) {
			c.Ingress.TLS.Production = true
			for i := range c.Environments {
				c.Environments[i].Ingress.TLS.Production = true
			}
		} else {
			c.Ingress.TLS.Production = false
			for i := range c.Environments {
				c.Environments[i].Ingress.TLS.Production = false
			}
		}
	}

	// StorageConfig is reused between multiple storage configuration types and needs to be handled explicitly
	if os.Getenv(RequirementStorageBackupURL) != "" {
		c.Storage = append(c.Storage, StorageConfig{
			Name: backupName,
			URL:  os.Getenv(RequirementStorageBackupURL),
		})
	}

	if os.Getenv(RequirementStorageLogsURL) != "" {
		c.Storage = append(c.Storage, StorageConfig{
			Name: logsName,
			URL:  os.Getenv(RequirementStorageLogsURL),
		})
	}

	if os.Getenv(RequirementStorageReportsURL) != "" {
		c.Storage = append(c.Storage, StorageConfig{
			Name: reportsName,
			URL:  os.Getenv(RequirementStorageReportsURL),
		})
	}

	if os.Getenv(RequirementStorageRepositoryURL) != "" {
		c.Storage = append(c.Storage, StorageConfig{
			Name: repositoryName,
			URL:  os.Getenv(RequirementStorageRepositoryURL),
		})
	}

	if os.Getenv(RequirementDevEnvApprovers) != "" {
		rawApprovers := os.Getenv(RequirementDevEnvApprovers)
		for _, approver := range strings.Split(rawApprovers, ",") {
			c.Cluster.DevEnvApprovers = append(c.Cluster.DevEnvApprovers, strings.TrimSpace(approver))
		}
	}

	// set this if its not currently configured
	if c.Cluster.Provider == "gke" {
		if c.Cluster.GKEConfig == nil {
			c.Cluster.GKEConfig = &GKEConfig{}
		}

		if c.Cluster.GKEConfig.ProjectNumber == "" {
			if gkeProjectNumber != nil {
				projectNumber, err := gkeProjectNumber(c.Cluster.ProjectID)
				if err != nil {
					log.Logger().Warnf("unable to determine gke project number - %s", err)
				}
				c.Cluster.GKEConfig.ProjectNumber = projectNumber
			}
		}
	}
}

func envVarBoolean(value string) bool {
	return value == constTrue || value == "yes"
}

func (c *RequirementsConfig) GetStorageURL(name string) string {
	for _, s := range c.Storage {
		if s.Name == name {
			return s.URL
		}
	}
	// default to empty string if no matching storage found for name
	return ""
}

func (c *RequirementsConfig) AddOrUpdateStorageURL(name, storageURL string) {
	found := false
	for i, s := range c.Storage {
		if s.Name == name {
			c.Storage[i].URL = storageURL
			found = true
		}
	}
	if !found {
		c.Storage = append(c.Storage, StorageConfig{
			Name: name,
			URL:  storageURL,
		})
	}
}

func (c *RequirementsConfig) RemoveStorageURL(name string) {
	for i, s := range c.Storage {
		if s.Name == name {
			c.Storage[i] = c.Storage[len(c.Storage)-1]    // Copy last element to index i.
			c.Storage[len(c.Storage)-1] = StorageConfig{} // Erase last element (write zero value).
			c.Storage = c.Storage[:len(c.Storage)-1]
			break
		}
	}
}
