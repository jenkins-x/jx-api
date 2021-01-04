package validation

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestOpenAPIJxRequirementsValidation(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/happy/jx-requirements.yml")
	assert.NoError(t, err)
	var requirements map[string]interface{}
	err = yaml.Unmarshal(data, &requirements)
	assert.NoError(t, err)
	fieldList, err := ValidateCustomResource("github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1.Requirements", requirements)
	assert.NoError(t, err)
	assert.Nil(t, fieldList)
}

func TestOpenAPIJxRequirementsValidationFailure(t *testing.T) {
	data, err := ioutil.ReadFile("test_data/sad/jx-requirements.yml")
	assert.NoError(t, err)
	var requirements map[string]interface{}
	err = yaml.Unmarshal(data, &requirements)

	assert.NoError(t, err)
	fieldList, err := ValidateCustomResource("github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1.Requirements", requirements)
	assert.NoError(t, err)
	assert.NotNil(t, fieldList)
}
