//nolint:dupl
package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	v1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	testUser = &v1.User{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: v1.UserDetails{},
	}
)

func TestPatchUpdateUserNoModification(t *testing.T) {
	json, err := json.Marshal(testUser)
	if err != nil {
		assert.Failf(t, "unable to marshal test instance: %s", err.Error())
	}
	get := func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Header: defaultHeaders(), Body: bytesBody(json)}, nil
	}

	patch := func(*http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Header: defaultHeaders(), Body: bytesBody(json)}, nil
	}

	fakeClient := newClientForTest(get, patch)

	users := users{
		client: fakeClient,
		ns:     "default",
	}

	updated, err := users.PatchUpdate(testUser)
	assert.NoError(t, err)
	assert.Equal(t, testUser, updated)
}

func TestPatchUpdateUserWithChange(t *testing.T) {
	clonedUser := testUser.DeepCopy()
	clonedUser.Spec.Name = FactNameUpdate

	get := func(*http.Request) (*http.Response, error) {
		json, err := json.Marshal(testUser)
		if err != nil {
			assert.Failf(t, "unable to marshal test instance: %s", err.Error())
		}
		return &http.Response{StatusCode: 200, Header: defaultHeaders(), Body: bytesBody(json)}, nil
	}

	patch := func(*http.Request) (*http.Response, error) {
		json, err := json.Marshal(clonedUser)
		if err != nil {
			assert.Failf(t, "unable to marshal test instance: %s", err.Error())
		}
		return &http.Response{StatusCode: 200, Header: defaultHeaders(), Body: bytesBody(json)}, nil
	}

	fakeClient := newClientForTest(get, patch)

	users := users{
		client: fakeClient,
		ns:     "default",
	}

	updated, err := users.PatchUpdate(clonedUser)
	assert.NoError(t, err)
	assert.NotEqual(t, testUser, updated)
	assert.Equal(t, FactNameUpdate, updated.Spec.Name)
}

func TestPatchUpdateUserWithErrorInGet(t *testing.T) {
	get := func(*http.Request) (*http.Response, error) {
		return nil, errors.New(errorDuringGETMessage)
	}

	fakeClient := newClientForTest(get, nil)

	users := users{
		client: fakeClient,
		ns:     "default",
	}

	updated, err := users.PatchUpdate(testUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errorDuringGETMessage)
	assert.Nil(t, updated)
}

func TestPatchUpdateUserWithErrorInPatch(t *testing.T) {
	get := func(*http.Request) (*http.Response, error) {
		json, err := json.Marshal(testUser)
		if err != nil {
			assert.Failf(t, "unable to marshal test instance: %s", err.Error())
		}
		return &http.Response{StatusCode: 200, Header: defaultHeaders(), Body: bytesBody(json)}, nil
	}

	patch := func(*http.Request) (*http.Response, error) {
		return nil, errors.New(errorDuringPATCHMessage)
	}

	fakeClient := newClientForTest(get, patch)

	users := users{
		client: fakeClient,
		ns:     "default",
	}
	name := "susfu"
	clonedUser := testUser.DeepCopy()
	clonedUser.Spec.Name = name
	updated, err := users.PatchUpdate(clonedUser)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errorDuringPATCHMessage)
	assert.Nil(t, updated)
}
