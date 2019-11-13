// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"fmt"
	"github.com/CanonicalLtd/imagebuild/domain"
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

// SetAuthorizationHeader mocks setting the oauth headers
func (c *MockAuthClient) SetAuthorizationHeader(header http.Header, credentials *oauth.Credentials, method string, u *url.URL, form url.Values) error {
	return nil
}

// MockClient mocks the Launchpad build API
type MockClient struct{}

// Build mocks the build submission method
func (cli *MockClient) Build(bld *domain.BuildRequest) (string, error) {
	if bld.BoardID == "error" {
		return "", fmt.Errorf("MOCK error building image")
	}
	return "", nil
}

// GetLiveFSBuild mocks the live build record retrieval
func (cli *MockClient) GetLiveFSBuild(rawURL string) (*domain.LiveFSBuild, error) {
	if rawURL == "error" {
		return nil, fmt.Errorf("MOCK error fetching build")
	}
	return &domain.LiveFSBuild{}, nil
}

// BoardMeta mocks the board metadata retrieval
func (cli *MockClient) BoardMeta(boardID, osID string) domain.Metadata {
	return domain.Metadata{}
}

// mockDoRequest mocks performing an HTTP request
func mockDoRequest(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 201, Body: ioutil.NopCloser(strings.NewReader(""))}, nil
}

// mockDoRequestError mocks performing an HTTP request with an error
func mockDoRequestError(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: 401, Body: ioutil.NopCloser(strings.NewReader(""))}, nil
}
