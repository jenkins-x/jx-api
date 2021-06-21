module github.com/jenkins-x/jx-api/v4

go 1.15

require (
	github.com/docker/spdystream v0.0.0-20160310174837-449fdfce4d96 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/imdario/mergo v0.3.9
	github.com/jenkins-x/jx-logging/v3 v3.0.6
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/pkg/errors v0.9.1
	github.com/rawlingsj/jsonschema v0.0.0-20210511142122-a9c2cfdb7dcf // use a fork until https://github.com/alecthomas/jsonschema/issues/65 is fixed
	github.com/stretchr/testify v1.6.1
	github.com/vrischmann/envconfig v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v0.21.2
	k8s.io/code-generator v0.20.2
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13

)
