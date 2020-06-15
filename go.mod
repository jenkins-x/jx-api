module github.com/jenkins-x/jx-api

go 1.13

require (
	github.com/alecthomas/jsonschema v0.0.0-20200530073317-71f438968921
	github.com/cenkalti/backoff v2.1.1+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/spec v0.19.7
	github.com/imdario/mergo v0.3.8
	github.com/jenkins-x/jx-logging v0.0.8
	github.com/jenkins-x/jx/v2 v2.1.65
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.1-0.20180103174451-36e9d2ebbde5
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v0.0.5
	github.com/stoewer/go-strcase v1.2.0
	github.com/stretchr/testify v1.6.0
	github.com/vrischmann/envconfig v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/tools v0.0.0-20200415034506-5d8e1897c761
	k8s.io/api v0.16.5
	k8s.io/apimachinery v0.16.5
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
)
