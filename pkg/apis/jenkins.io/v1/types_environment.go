package v1

import (
	"encoding/json"
	"fmt"

	"github.com/jenkins-x/jx-logging/v3/pkg/log"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Namespace",type="string",JSONPath=".spec.namespace",description="The namespace used for the environment"
// +kubebuilder:printcolumn:name="Kind",type="string",JSONPath=".spec.kind",description="The kind of environment"
// +kubebuilder:printcolumn:name="Promotion",type="string",JSONPath=".spec.promotionStrategy",description="The strategy used for promoting to this environment"
// +kubebuilder:printcolumn:name="Order",type="integer",JSONPath=".spec.order",description="The order in which environments are automatically promoted"
// +kubebuilder:printcolumn:name="Git URL",type="string",JSONPath=".spec.source.url",description="The Git repository URL for the source of the environment configuration"
// +kubebuilder:printcolumn:name="Git Branch",type="string",JSONPath=".spec.source.ref",description="The git branch for the source of the environment configuration"
// +kubebuilder:resource:categories=all,shortName=env
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Environment represents an environment like Dev, Test, Staging, Production where code lives
type Environment struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec   EnvironmentSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status EnvironmentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// EnvironmentSpec is the specification of an Environment
type EnvironmentSpec struct {
	Label             string                `json:"label,omitempty" protobuf:"bytes,1,opt,name=label"`
	Namespace         string                `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
	Cluster           string                `json:"cluster,omitempty" protobuf:"bytes,3,opt,name=cluster"`
	PromotionStrategy PromotionStrategyType `json:"promotionStrategy,omitempty" protobuf:"bytes,4,opt,name=promotionStrategy"`
	Source            EnvironmentRepository `json:"source,omitempty" protobuf:"bytes,5,opt,name=source"`
	Order             int32                 `json:"order,omitempty" protobuf:"bytes,6,opt,name=order"`
	Kind              EnvironmentKindType   `json:"kind,omitempty" protobuf:"bytes,7,opt,name=kind"`
	PullRequestURL    string                `json:"pullRequestURL,omitempty" protobuf:"bytes,8,opt,name=pullRequestURL"`
	TeamSettings      TeamSettings          `json:"teamSettings,omitempty" protobuf:"bytes,9,opt,name=teamSettings"`
	PreviewGitSpec    PreviewGitSpec        `json:"previewGitInfo,omitempty" protobuf:"bytes,10,opt,name=previewGitInfo"`
	WebHookEngine     WebHookEngineType     `json:"webHookEngine,omitempty" protobuf:"bytes,11,opt,name=webHookEngine"`

	// RemoteCluster flag indicates if the Environment is deployed in a separate cluster to the Development Environment
	RemoteCluster bool `json:"remoteCluster,omitempty" protobuf:"bytes,12,opt,name=remoteCluster"`
}

// EnvironmentStatus is the status for an Environment resource
type EnvironmentStatus struct {
	Version string `json:"version,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvironmentList is a list of TypeMeta resources
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Environment `json:"items"`
}

// PromotionStrategyType is the type of a promotion strategy
type PromotionStrategyType string

const (
	// PromotionStrategyTypeManual specifies that promotion happens manually
	PromotionStrategyTypeManual PromotionStrategyType = "Manual"
	// PromotionStrategyTypeAutomatic specifies that promotion happens automatically
	PromotionStrategyTypeAutomatic PromotionStrategyType = "Auto"
	// PromotionStrategyTypeNever specifies that promotion is disabled for this environment
	PromotionStrategyTypeNever PromotionStrategyType = "Never"
)

// EnvironmentKindType is the kind of an environment
type EnvironmentKindType string

const (
	// EnvironmentKindTypePermanent specifies that the environment is a regular permanent one
	EnvironmentKindTypePermanent EnvironmentKindType = "Permanent"
	// EnvironmentKindTypePreview specifies that an environment is a Preview environment that lasts as long as a Pull Request
	EnvironmentKindTypePreview EnvironmentKindType = "Preview"
	// EnvironmentKindTypeTest specifies that an environment is a temporary one for a test
	EnvironmentKindTypeTest EnvironmentKindType = "Test"
	// EnvironmentKindTypeEdit specifies that an environment is a developers editing workspace
	EnvironmentKindTypeEdit EnvironmentKindType = "Edit"
	// EnvironmentKindTypeDevelopment specifies that an environment is a development environment; for developer tools like Jenkins, Nexus etc
	EnvironmentKindTypeDevelopment EnvironmentKindType = "Development"
)

// PromotionEngineType is the type of promotion implementation the team uses
type PromotionEngineType string

// ProwEngineType is the type of prow execution engine
type ProwEngineType string

// ProwConfigType is the type of prow configuration
type ProwConfigType string

const (
	// ProwConfigScheduler when we use the Scheduler CRDs to generate the Prow ConfigMaps
	ProwConfigScheduler ProwConfigType = "Scheduler"
)

// WebHookEngineType is the type of webhook processing implementation the team uses
type WebHookEngineType string

const (
	// WebHookEngineLighthouse specifies that we use lighthouse for webhooks
	// see: https://github.com/jenkins-x/lighthouse
	WebHookEngineLighthouse WebHookEngineType = "Lighthouse"
)

// IsPermanent returns true if this environment is permanent
func (e EnvironmentKindType) IsPermanent() bool {
	switch e {
	case EnvironmentKindTypePreview, EnvironmentKindTypeTest, EnvironmentKindTypeEdit:
		return false
	default:
		return true
	}
}

// PromotionStrategyTypeValues is the list of all values
var PromotionStrategyTypeValues = []string{
	string(PromotionStrategyTypeAutomatic),
	string(PromotionStrategyTypeManual),
	string(PromotionStrategyTypeNever),
}

// EnvironmentRepositoryType is the repository type
type EnvironmentRepositoryType string

const (
	// EnvironmentRepositoryTypeGit specifies that a git repository is used
	EnvironmentRepositoryTypeGit EnvironmentRepositoryType = "Git"
)

// EnvironmentRepository is the repository for an environment using GitOps
type EnvironmentRepository struct {
	Kind EnvironmentRepositoryType `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	URL  string                    `json:"url,omitempty" protobuf:"bytes,2,opt,name=url"`
	Ref  string                    `json:"ref,omitempty" protobuf:"bytes,3,opt,name=ref"`
}

// DeployOptions configures options for how to deploy applications by default such as using progressive delivery or using horizontal pod autoscaler
type DeployOptions struct {
	// Canary should we enable canary rollouts (progressive delivery) for apps by default
	Canary bool `json:"canary,omitempty" protobuf:"bytes,1,opt,name=canary"`

	// should we use the horizontal pod autoscaler on new apps by default?
	HPA bool `json:"hpa,omitempty" protobuf:"bytes,2,opt,name=hpa"`
}

// TeamSettings the default settings for a team
type TeamSettings struct {
	BranchPatterns      string               `json:"branchPatterns,omitempty" protobuf:"bytes,3,opt,name=branchPatterns"`
	ForkBranchPatterns  string               `json:"forkBranchPatterns,omitempty" protobuf:"bytes,4,opt,name=forkBranchPatterns"`
	QuickstartLocations []QuickStartLocation `json:"quickstartLocations,omitempty" protobuf:"bytes,5,opt,name=quickstartLocations"`
	PostPreviewJobs     []batchv1.Job        `json:"postPreviewJobs,omitempty" protobuf:"bytes,9,opt,name=postPreviewJobs"`
	PromotionEngine     PromotionEngineType  `json:"promotionEngine,omitempty" protobuf:"bytes,10,opt,name=promotionEngine"`
	GitServer           string               `json:"gitServer,omitempty" protobuf:"bytes,13,opt,name=gitServer" command:"gitserver" commandUsage:"Default git server for new repositories"`
	Organisation        string               `json:"organisation,omitempty" protobuf:"bytes,14,opt,name=organisation" command:"organisation" commandUsage:"Default git organisation for new repositories"`
	EnvOrganisation     string               `json:"envOrganisation,omitempty" protobuf:"bytes,14,opt,name=envOrganisation" command:"envOrganisation" commandUsage:"Default git organisation for new environment repositories"`
	PipelineUsername    string               `json:"pipelineUsername,omitempty" protobuf:"bytes,15,opt,name=pipelineUsername" command:"pipelineusername" commandUsage:"User used by pipeline. Is given write permission on new repositories."`
	PipelineUserEmail   string               `json:"pipelineUserEmail,omitempty" protobuf:"bytes,15,opt,name=pipelineUserEmail" command:"pipelineuseremail" commandUsage:"Users email used by pipeline. Is given write permission on new repositories."`
	DockerRegistryOrg   string               `json:"dockerRegistryOrg,omitempty" protobuf:"bytes,16,opt,name=dockerRegistryOrg" command:"dockerregistryorg" commandUsage:"Docker registry organisation used for new projects in Jenkins X."`
	GitPublic           bool                 `json:"gitPublic,omitempty" protobuf:"bytes,17,opt,name=gitPublic" command:"gitpublic" commandUsage:"Are new repositories public by default"`
	KubeProvider        string               `json:"kubeProvider,omitempty" protobuf:"bytes,18,opt,name=kubeProvider"`
	AppsRepository      string               `json:"appsRepository,omitempty" protobuf:"bytes,19,opt,name=appsRepository"`
	StorageLocations    []StorageLocation    `json:"storageLocations,omitempty" protobuf:"bytes,21,opt,name=storageLocations"`

	// DeployKind what kind of deployment ("default" uses regular Kubernetes Services and Deployments, "knative" uses the Knative Service resource instead)
	DeployKind string `json:"deployKind,omitempty" protobuf:"bytes,24,opt,name=deployKind"`

	// AppsPrefixes is the list of prefixes for appNames
	AppsPrefixes     []string          `json:"appPrefixes,omitempty" protobuf:"bytes,27,opt,name=appPrefixes"`
	DefaultScheduler ResourceReference `json:"defaultScheduler,omitempty" protobuf:"bytes,28,opt,name=defaultScheduler"`

	// ProwConfig is the way we manage prow configurations
	ProwConfig ProwConfigType `json:"prowConfig,omitempty" protobuf:"bytes,29,opt,name=prowConfig"`

	// BootRequirements is a marshaled string of the jx-requirements.yml used in the most recent run for this cluster
	BootRequirements string `json:"bootRequirements,omitempty" protobuf:"bytes,31,opt,name=bootRequirements"`

	// DeployOptions configures options for how to deploy applications by default such as using canary rollouts (progressive delivery) or using horizontal pod autoscaler
	DeployOptions *DeployOptions `json:"deployOptions,omitempty" protobuf:"bytes,32,opt,name=deployOptions"`
}

// StorageLocation
type StorageLocation struct {
	Classifier string `json:"classifier,omitempty" protobuf:"bytes,1,opt,name=classifier"`
	GitURL     string `json:"gitUrl,omitempty" protobuf:"bytes,2,opt,name=gitUrl"`
	GitBranch  string `json:"gitBranch,omitempty" protobuf:"bytes,3,opt,name=gitBranch"`
	BucketURL  string `json:"bucketUrl,omitempty" protobuf:"bytes,4,opt,name=bucketUrl"`
}

// QuickStartLocation
type QuickStartLocation struct {
	GitURL   string   `json:"gitUrl,omitempty" protobuf:"bytes,1,opt,name=gitUrl"`
	GitKind  string   `json:"gitKind,omitempty" protobuf:"bytes,2,opt,name=gitKind"`
	Owner    string   `json:"owner,omitempty" protobuf:"bytes,3,opt,name=owner"`
	Includes []string `json:"includes,omitempty" protobuf:"bytes,4,opt,name=includes"`
	Excludes []string `json:"excludes,omitempty" protobuf:"bytes,5,opt,name=excludes"`
}

// PreviewGitSpec is the preview git branch/pull request details
type PreviewGitSpec struct {
	Name            string   `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	URL             string   `json:"url,omitempty" protobuf:"bytes,2,opt,name=url"`
	User            UserSpec `json:"user,omitempty" protobuf:"bytes,3,opt,name=user"`
	Title           string   `json:"title,omitempty" protobuf:"bytes,4,opt,name=title"`
	Description     string   `json:"description,omitempty" protobuf:"bytes,5,opt,name=description"`
	BuildStatus     string   `json:"buildStatus,omitempty" protobuf:"bytes,6,opt,name=buildStatus"`
	BuildStatusURL  string   `json:"buildStatusUrl,omitempty" protobuf:"bytes,7,opt,name=buildStatusUrl"`
	ApplicationName string   `json:"appName,omitempty" protobuf:"bytes,8,opt,name=appName"`
	ApplicationURL  string   `json:"applicationURL,omitempty" protobuf:"bytes,9,opt,name=applicationURL"`
}

// UserSpec is the user details
type UserSpec struct {
	Username string `json:"username,omitempty" protobuf:"bytes,1,opt,name=username"`
	Name     string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	LinkURL  string `json:"linkUrl,omitempty" protobuf:"bytes,3,opt,name=linkUrl"`
	ImageURL string `json:"imageUrl,omitempty" protobuf:"bytes,4,opt,name=imageUrl"`
}

// EnvironmentFilter specifies the environments to apply the role binding to
type EnvironmentFilter struct {
	Kind     EnvironmentKindType `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	Includes []string            `json:"includes,omitempty" protobuf:"bytes,2,opt,name=includes"`
	Excludes []string            `json:"excludes,omitempty" protobuf:"bytes,3,opt,name=excludes"`
}

// EnvironmentRoleBindingStatus is the status for an EnvironmentRoleBinding resource
type EnvironmentRoleBindingStatus struct {
	Version string `json:"version,omitempty"`
}

// StorageLocationOrDefault returns the storage location if there is one or returns the default storage configuration
func (t *TeamSettings) StorageLocationOrDefault(classifier string) StorageLocation {
	for idx, sl := range t.StorageLocations {
		if sl.Classifier == classifier {
			return t.StorageLocations[idx]
		}
	}
	return t.StorageLocation("default")
}

// StorageLocation returns the storage location, lazily creating one if one does not already exist
func (t *TeamSettings) StorageLocation(classifier string) StorageLocation {
	for idx, sl := range t.StorageLocations {
		if sl.Classifier == classifier {
			return t.StorageLocations[idx]
		}
	}
	newStorageLocation := StorageLocation{
		Classifier: classifier,
	}
	t.SetStorageLocation(classifier, newStorageLocation)
	return newStorageLocation
}

// SetStorageLocation stores the given storage location in the team settings
func (t *TeamSettings) SetStorageLocation(classifier string, storage StorageLocation) {
	storage.Classifier = classifier
	for idx, sl := range t.StorageLocations {
		if sl.Classifier == classifier {
			t.StorageLocations[idx] = storage
			return
		}
	}
	t.StorageLocations = append(t.StorageLocations, storage)
}

// GetDeployOptions returns the default deploy options for a team
func (t *TeamSettings) GetDeployOptions() DeployOptions {
	if t.DeployOptions != nil {
		return *t.DeployOptions
	}
	return DeployOptions{}
}

// UnmarshalJSON method handles the rename of GitPrivate to GitPublic.
func (t *TeamSettings) UnmarshalJSON(data []byte) error {
	// need a type alias to go into infinite loop
	type Alias TeamSettings
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

	_, gitPublicSet := raw["gitPublic"]
	private, gitPrivateSet := raw["gitPrivate"]

	if gitPrivateSet && gitPublicSet {
		return fmt.Errorf("found settings for GitPublic as well as GitPrivate in TeamSettings, only GitPublic should be used")
	}

	if gitPrivateSet {
		log.Logger().Debug("GitPrivate specified in TeamSettings. GitPrivate is deprecated use GitPublic instead.")
		privateString := string(private)
		if privateString == "true" {
			t.GitPublic = false
		} else {
			t.GitPublic = true
		}
	}
	return nil
}

// IsLighthouse returns true if we are using lighthouse as the webhook handler
func (e *EnvironmentSpec) IsLighthouse() bool {
	return e.WebHookEngine == WebHookEngineLighthouse
}

// IsEmpty returns true if the storage location is empty
func (s *StorageLocation) IsEmpty() bool {
	return s.GitURL == "" && s.BucketURL == ""
}

// Description returns the textual description of the storage location
func (s *StorageLocation) Description() string {
	if s.GitURL != "" {
		return s.GitURL + " branch: " + s.GetGitBranch()
	}
	if s.BucketURL != "" {
		return s.BucketURL
	}
	return "current git repo"
}

// GetGitBranch returns the git branch to use when using git storage
func (s *StorageLocation) GetGitBranch() string {
	branch := s.GitBranch
	if branch == "" {
		branch = "gh-pages"
	}
	return branch
}
