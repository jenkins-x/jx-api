package util

import (
	"net/url"
	"strings"
)

// SanitizeURL sanitizes by stripping the user and password
func SanitizeURL(unsanitizedUrl string) string {
	u, err := url.Parse(unsanitizedUrl)
	if err != nil {
		return unsanitizedUrl
	}
	return stripCredentialsFromURL(u)
}

// stripCredentialsFromURL strip credentials from URL
func stripCredentialsFromURL(u *url.URL) string {
	pass, hasPassword := u.User.Password()
	userName := u.User.Username()
	if hasPassword {
		textToReplace := pass + "@"
		textToReplace = ":" + textToReplace
		if userName != "" {
			textToReplace = userName + textToReplace
		}
		return strings.Replace(u.String(), textToReplace, "", 1)
	}
	return u.String()
}
