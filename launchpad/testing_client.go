// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"github.com/gomodule/oauth1/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// MockAuthClient mocks the oauth client
type MockAuthClient struct {
	body string
}

// Post issues a POST with the specified form.
func (c *MockAuthClient) Post(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error) {
	return &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(c.body)),
	}, nil
}

// MockClient mocks the Launchpad build API
type MockClient struct{}

// Build mocks the build submission method
func (cli *MockClient) Build(boardID, osID string) error {
	return nil
}
