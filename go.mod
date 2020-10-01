module github.com/jenkins-x/jx-api/v3

go 1.15

require (
	github.com/alecthomas/jsonschema v0.0.0-20190504002508-159cbd5dba26
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/imdario/mergo v0.3.9
	github.com/jenkins-x/jx-logging/v3 v3.0.0
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/pkg/errors v0.9.1
	github.com/stoewer/go-strcase v1.2.0
	github.com/stretchr/testify v1.6.1
	github.com/vrischmann/envconfig v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	k8s.io/api v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	k8s.io/code-generator v0.19.2
	k8s.io/kube-openapi v0.0.0-20200923155610-8b5066479488 // indirect
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
)
