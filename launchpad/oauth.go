// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/gomodule/oauth1/oauth"
	"net/http"
	"net/url"
)

// OAuthClient defines the interface for the oauth methods we need
type OAuthClient interface {
	//Get(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error)
	Post(client *http.Client, credentials *oauth.Credentials, urlStr string, form url.Values) (*http.Response, error)
}

// NewOAuthClient creates an implementation of a standard oauth1 client
func NewOAuthClient(settings *config.Settings) *oauth.Client {
	return &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  settings.LPToken,
			Secret: settings.LPSecret,
		},
	}
}
