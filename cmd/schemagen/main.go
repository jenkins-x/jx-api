package main

import (
	"github.com/jenkins-x/jx-api/v4/pkg/apis/core"
	"github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1"
	jenkins_io "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io"
	"github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx-api/v4/pkg/schemagen"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"os"
)

var (
	resourceKinds = []schemagen.ResourceKind{
		{
			APIVersion: core.GroupAndVersion,
			Name:       "requirements",
			Resource:   &v4beta1.Requirements{},
		},
		{
			APIVersion: jenkins_io.GroupAndVersion,
			Name:       "environment",
			Resource:   &v1.Environment{},
		},
		{
			APIVersion: jenkins_io.GroupAndVersion,
			Name:       "source-repository",
			Resource:   &v1.SourceRepository{},
		},
		{
			APIVersion: jenkins_io.GroupAndVersion,
			Name:       "pipeline-activity",
			Resource:   &v1.PipelineActivity{},
		},
		{
			APIVersion: jenkins_io.GroupAndVersion,
			Name:       "release",
			Resource:   &v1.Release{},
		},
	}
)

func main() {
	out := "schema"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	err := schemagen.GenerateSchemas(resourceKinds, out)
	if err != nil {
		log.Logger().Errorf("failed: %v", err)
		os.Exit(1)
	}
	log.Logger().Infof("completed the plugin generator")
	os.Exit(0)
}
