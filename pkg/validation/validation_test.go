package validation

import (
	"io/ioutil"
	"testing"

	"github.com/jenkins-x/jx-api/v4/pkg/apis/core"
	"github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidationSuccess(t *testing.T) {
	crdFileName := v4beta1.CrdFilename
	requirementsData, err := ioutil.ReadFile("test_data/happy/jx-requirements.yml")
	assert.NoError(t, err)
	var requirements map[string]interface{}
	err = yaml.Unmarshal(requirementsData, &requirements)
	validationList, err := ValidateCrd(crdFileName, core.Version, requirements)
	assert.NoError(t, err)
	assert.Nil(t, validationList)
}

func TestValidationFailure(t *testing.T) {
	crdFileName := v4beta1.CrdFilename
	requirementsData, err := ioutil.ReadFile("test_data/sad/jx-requirements.yml")
	assert.NoError(t, err)
	var requirements map[string]interface{}
	err = yaml.Unmarshal(requirementsData, &requirements)
	validationList, err := ValidateCrd(crdFileName, core.Version, requirements)
	assert.NoError(t, err)
	assert.NotNil(t, validationList)
	assert.Equal(t, field.ErrorTypeRequired, validationList[0].Type)
	assert.Equal(t, "[].spec.ingress", validationList[0].Field)
}
