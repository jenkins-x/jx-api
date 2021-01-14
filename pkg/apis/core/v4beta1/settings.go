package v4beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// SettingsFileName the default file name of the settings file
	SettingsFileName = "settings.yaml"
)

// Settings represents application specific settings for use inside a pipeline of an application
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
type Settings struct {
	metav1.TypeMeta `json:",inline"`

	// Spec the definition of the settings
	Spec SettingsConfig `json:"spec"`
}

// SettingsConfig contains the optional overrides you can specify on a per application basis
type SettingsConfig struct {
	// GitURL the git URL for your development cluster where the default environments and cluster configuration are specified
	GitURL string `json:"gitUrl,omitempty"`
	// Destination settings to define where release artifacts go in terms of containers and charts
	Destination *DestinationConfig `json:"destination"`
	// PromoteEnvironments the environments for promotion
	PromoteEnvironments []EnvironmentConfig `json:"promoteEnvironments,omitempty"`
	// IgnoreDevEnvironments if enabled do not inherit any environments from the
	IgnoreDevEnvironments bool `json:"ignoreDevEnvironments,omitempty"`
}
