// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"fmt"
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/domain"
	"github.com/gomodule/oauth1/oauth"
	"io/ioutil"
	"net/url"
)

const apiURL = "https://api.launchpad.net/devel"

// BuildClient defines the interface for the Launchpad client for building images
type BuildClient interface {
	Build(img *domain.BuildRequest) (string, error)
	GetLiveFSBuild(urlString string) (*domain.LiveFSBuild, error)
	BoardMeta(boardID, osID string) domain.Metadata
}

// Client defines a Launchpad client
type Client struct {
	settings    *config.Settings
	oauthClient OAuthClient
	credentials oauth.Credentials
}

// NewClient creates a new Launchpad client
func NewClient(settings *config.Settings, authClient OAuthClient) (*Client, error) {
	client := &Client{
		settings:    settings,
		oauthClient: authClient,
		credentials: oauth.Credentials{
			Token:  settings.LPToken,
			Secret: settings.LPSecret,
		},
	}

	// Read the metadata for the board configurations
	err := client.readMetadata(settings.BuildsPath)

	return client, err
}

// Build starts a build
func (cli *Client) Build(img *domain.BuildRequest) (string, error) {
	// Get the metadata for the selection
	distroArchSeries, liveFS, meta, err := cli.buildMetadata(img)
	if err != nil {
		return "", err
	}

	// Start the build
	return cli.requestBuild(distroArchSeries, liveFS, meta)
}

// GetLiveFSBuild fetches a build's details
func (cli *Client) GetLiveFSBuild(rawURL string) (*domain.LiveFSBuild, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	// Call the API
	resp, err := cli.httpDo("GET", u, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return nil, fmt.Errorf("error submitting %d build: %v", resp.StatusCode, string(bodyBytes))
	}

	// Parse the response
	liveBuild, err := decodeLiveFSBuild(resp)
	if err != nil {
		return liveBuild, err
	}

	// Get the download URL if the build is completed successfully
	if liveBuild.State == SuccessfullyBuilt {
		err = cli.getBuildURL(u, liveBuild)
	}
	return liveBuild, err
}

// BoardMeta fetches the metadata for the board and OS
func (cli *Client) BoardMeta(boardID, osID string) domain.Metadata {
	// Get the metadata for the board and OS
	key := fmt.Sprintf("%s-%s", boardID, osID)
	return boards[key]
}
