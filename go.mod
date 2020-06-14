module github.com/jenkins-x/jx-client

go 1.13

require (
	github.com/alecthomas/jsonschema v0.0.0-20200530073317-71f438968921
	github.com/ghodss/yaml v1.0.0
	github.com/imdario/mergo v0.3.8
	github.com/jenkins-x/jx-logging v0.0.8
	github.com/pkg/errors v0.9.1
	github.com/stoewer/go-strcase v1.2.0
	github.com/stretchr/testify v1.6.0
	github.com/vrischmann/envconfig v1.2.0
	github.com/xeipuuv/gojsonschema v1.2.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451 // indirect
)

replace k8s.io/api => k8s.io/api v0.16.5

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20190819143841-305e1cef1ab1

replace k8s.io/apimachinery => k8s.io/apimachinery v0.16.5

replace k8s.io/client-go => k8s.io/client-go v0.16.5

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190819143637-0dbe462fe92d
