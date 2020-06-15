package util_test

import (
	"github.com/jenkins-x/jx-client/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSanitizeURL(t *testing.T) {
	t.Parallel()
	tests := map[string]string{
		"http://test.com":                 "http://test.com",
		"http://user:test@github.com":     "http://github.com",
		"https://user:test@github.com":    "https://github.com",
		"https://user:@github.com":        "https://github.com",
		"https://:pass@github.com":        "https://github.com",
		"git@github.com:jenkins-x/jx-client.git": "git@github.com:jenkins-x/jx-client.git",
		"invalid/url":                     "invalid/url",
	}

	for test, expected := range tests {
		t.Run(test, func(t *testing.T) {
			actual := util.SanitizeURL(test)
			assert.Equal(t, expected, actual, "for url: %s", test)
		})
	}
}
