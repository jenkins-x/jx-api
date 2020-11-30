package util_test

import (
	"testing"

	"github.com/jenkins-x/jx-api/v4/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeURL(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"http://test.com":                     "http://test.com",
		"http://user:test@github.com":         "http://github.com",
		"https://user:test@github.com":        "https://github.com",
		"https://user:@github.com":            "https://github.com",
		"https://:pass@github.com":            "https://github.com",
		"git@github.com:jenkins-x/jx-api.git": "git@github.com:jenkins-x/jx-api.git",
		"invalid/url":                         "invalid/url",
	}

	for test, expected := range tests {
		t.Run(test, func(t *testing.T) {
			actual := util.SanitizeURL(test)
			assert.Equal(t, expected, actual, "for url: %s", test)
		})
	}
}
