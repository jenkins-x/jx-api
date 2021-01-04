package validation

import (
	"github.com/jenkins-x/jx-api/v4/pkg/generated/openapi"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateCustomResource(scheme string, resource interface{}) (field.ErrorList, error) {
	customResourceValidation := GetCustomResourceValidation(scheme, []GetAPIDefinitions{openapi.GetOpenAPIDefinitions})
	validator, _, err := validation.NewSchemaValidator(customResourceValidation)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create schema validator for custom resource")
	}

	return validation.ValidateCustomResource(nil, resource, validator), nil
}
