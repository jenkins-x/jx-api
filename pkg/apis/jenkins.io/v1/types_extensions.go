package v1

import (
	"github.com/jenkins-x/jx-api/v4/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="URL",type="string",JSONPath=".spec.url",description="The URL of the git repository"
// +kubebuilder:printcolumn:name="Description",type="string",JSONPath=".spec.description",description="A description of the source code repository - non-functional user-data"
// +kubebuilder:resource:categories=all,shortName=sourcerepo;srcrepo;sr
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
// SourceRepository is the metadata for an Application/Project/SourceRepository
type SourceRepository struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec SourceRepositorySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// Sanitize sanitizes the source repository URL
func (repo *SourceRepository) Sanitize() {
	repo.Spec.URL = util.SanitizeURL(repo.Spec.URL)
	repo.Spec.HTTPCloneURL = util.SanitizeURL(repo.Spec.HTTPCloneURL)
	repo.Spec.SSHCloneURL = util.SanitizeURL(repo.Spec.SSHCloneURL)
	// The URL is stored sometimes in the provider and provider name
	repo.Spec.Provider = util.SanitizeURL(repo.Spec.Provider)
	repo.Spec.ProviderName = util.SanitizeURL(repo.Spec.ProviderName)
}

// +kubebuilder:object:root=true

// SourceRepositoryList is a structure used by k8s to store lists of apps
type SourceRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SourceRepository `json:"items"`
}

// SourceRepositorySpec provides details of the metadata for an App
type SourceRepositorySpec struct {
	Description string `json:"description,omitempty" protobuf:"bytes,1,opt,name=description"`
	// Provider stores the URL of the git provider such as https://github.com
	Provider string `json:"provider,omitempty" protobuf:"bytes,2,opt,name=provider"`
	Org      string `json:"org,omitempty" protobuf:"bytes,3,opt,name=org"`
	Repo     string `json:"repo,omitempty" protobuf:"bytes,4,opt,name=repo"`
	// ProviderName is a logical name for the provider without any URL scheme which can be used in a label selector
	ProviderName string `json:"providerName,omitempty" protobuf:"bytes,5,opt,name=providerName"`
	// ProviderKind is the kind of provider (github / bitbucketcloud / bitbucketserver etc)
	ProviderKind string `json:"providerKind,omitempty" protobuf:"bytes,6,opt,name=providerKind"`
	// URL is the web URL of the project page
	URL string `json:"url,omitempty" protobuf:"bytes,7,opt,name=url"`
	// SSHCloneURL is the git URL to clone this repository using SSH
	SSHCloneURL string `json:"sshCloneURL,omitempty" protobuf:"bytes,8,opt,name=sshCloneURL"`
	// HTTPCloneURL is the git URL to clone this repository using HTTP/HTTPS
	HTTPCloneURL string `json:"httpCloneURL,omitempty" protobuf:"bytes,9,opt,name=httpCloneURL"`
	// Scheduler a reference to a custom scheduler otherwise we default to the Team's Scededuler
	Scheduler ResourceReference `json:"scheduler,omitempty" protobuf:"bytes,10,opt,name=scheduler"`
}

// PipelineExtension defines the image and command of an app which wants to modify/extend the pipeline
type PipelineExtension struct {
	// Name of the container specified as a DNS_LABEL.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// Docker image name.
	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`
	// Entrypoint array. Not executed within a shell.
	Command string `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`
	// Arguments to the entrypoint.
	Args []string `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`
}

// +kubebuilder:object:root=true

// PluginList is a list of Plugins available for a team
type PluginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Plugin `json:"items"`
}

// +genclient
// +genclient:noStatus
// +kubebuilder:object:root=true
// +k8s:openapi-gen=true

// Plugin represents a binary plugin installed into this Jenkins X team
type Plugin struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec PluginSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// PluginSpec provides details of a binary plugin available for a team
type PluginSpec struct {
	SubCommand  string   `json:"subCommand,omitempty"  protobuf:"bytes,3,opt,name=subCommand"`
	Group       string   `json:"group,omitempty"  protobuf:"bytes,4,opt,name=group"`
	Binaries    []Binary `json:"binaries,omitempty" protobuf:"bytes,7opt,name=binaries"`
	Description string   `json:"description,omitempty"  protobuf:"bytes,2,opt,name=description"`
	Name        string   `json:"name,omitempty"  protobuf:"bytes,5,opt,name=name"`
	Version     string   `json:"version,omitempty"  protobuf:"bytes,6,opt,name=version"`
}

// Binary provies the details of a downloadable binary
type Binary struct {
	Goarch string `json:"goarch,omitempty"  protobuf:"bytes,1,opt,name=goarch"`
	Goos   string `json:"goos,omitempty"  protobuf:"bytes,2,opt,name=goos"`
	URL    string `json:"url,omitempty"  protobuf:"bytes,3,opt,name=url"`
}
