package util

import (
	"github.com/ghodss/yaml"
	schemagen "github.com/invopop/jsonschema"
	"github.com/xeipuuv/gojsonschema"
	corev1 "k8s.io/api/core/v1"
)

// GenerateSchema generates a JSON schema for the given struct type and returns it.
func GenerateSchema(target interface{}, allowAdditionalProperties bool) *schemagen.Schema {
	reflector := schemagen.Reflector{
		IgnoredTypes: []interface{}{
			corev1.Container{},
		},
		AllowAdditionalProperties: allowAdditionalProperties,
		//ExpandedStruct: true,
		RequiredFromJSONSchemaTags: true,
	}
	return reflector.Reflect(target)
}

// ValidateYaml generates a JSON schema for the given struct type, and then validates the given YAML against that
// schema, ignoring Containers and missing fields.
func ValidateYaml(target interface{}, data []byte) ([]string, error) {
	return ValidateYamlLenient(target, data, false)
}

// ValidateYamlLenient generates a JSON schema for the given struct type, and then validates the given YAML against that
// schema, ignoring Containers and missing fields. If allowAdditionalProperties is true additional keys in JSON objects are allowed.
func ValidateYamlLenient(target interface{}, data []byte, allowAdditionalProperties bool) ([]string, error) {
	schema := GenerateSchema(target, allowAdditionalProperties)
	dataAsJSON, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	schemaLoader := gojsonschema.NewGoLoader(schema)
	documentLoader := gojsonschema.NewBytesLoader(dataAsJSON)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, err
	}
	if !result.Valid() {
		errMsgs := []string{}
		for _, e := range result.Errors() {
			errMsgs = append(errMsgs, e.String())
		}
		return errMsgs, nil
	}

	return nil, nil
}
