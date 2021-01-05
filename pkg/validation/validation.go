package validation

import (
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateCrd(crdFilename string, version string, resource map[string]interface{}) (field.ErrorList, error) {
	data, err := ioutil.ReadFile(filepath.Join("..", "..", "config", "crd", "bases", crdFilename))
	if err != nil {
		return nil, err
	}
	crd := apiextensions.CustomResourceDefinition{}

	err = yaml.Unmarshal(data, &crd)
	if err != nil {
		return nil, err
	}
	validator, _, err := validation.NewSchemaValidator(crd.Spec.Validation)
	return validation.ValidateCustomResource(&field.Path{}, resource, validator), nil
}
