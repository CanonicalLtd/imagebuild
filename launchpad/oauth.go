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
	SetAuthorizationHeader(header http.Header, credentials *oauth.Credentials, method string, u *url.URL, form url.Values) error
}

// NewOAuthClient creates an implementation of a standard oauth1 client
func NewOAuthClient(settings *config.Settings) *oauth.Client {
	return &oauth.Client{
		TemporaryCredentialRequestURI: "https://launchpad.net/+request-token",
		ResourceOwnerAuthorizationURI: "https://launchpad.net/+authorize-token",
		TokenRequestURI:               "https://launchpad.net/+access-token",
		Credentials: oauth.Credentials{
			Token: settings.LPConsumer, // Note that the consumer is used here, not the access token
		},
		SignatureMethod: oauth.PLAINTEXT,
	}
}
