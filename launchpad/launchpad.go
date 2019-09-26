// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/gomodule/oauth1/oauth"
	"log"
	"net/url"
	"path"
)

const apiURL = "https://api.launchpad.net/devel"

// BuildClient defines the interface for the Launchpad client for building images
type BuildClient interface {
	Build(boardID, osID string) error
}

// Client defines a Launchpad client
type Client struct {
	oauthClient OAuthClient
	credentials oauth.Credentials
}

// NewClient creates a new Launchpad client
func NewClient(settings *config.Settings, authClient OAuthClient) *Client {
	return &Client{
		oauthClient: authClient,
		credentials: oauth.Credentials{
			Token:  settings.LPToken,
			Secret: settings.LPSecret,
		},
	}
}

// Build starts a build
func (cli *Client) Build(boardID, osID string) error {
	// Get the metadata for the selection
	distroArchSeries, meta, err := cli.buildMetadata(boardID, osID)
	if err != nil {
		return err
	}

	// Start the build
	return cli.requestBuild(distroArchSeries, meta)
}

// requestBuild starts a new build as per https://launchpad.net/+apidoc/devel.html#livefs-requestBuild
func (cli *Client) requestBuild(das, metadata string) error {
	// Set the parameters
	params := map[string][]string{
		"ws.op":              {"requestBuild"},
		"pocket":             {"Release"},
		"archive":            {"https://api.launchpad.net/1.0/ubuntu/+archive/primary"},
		"distro_arch_series": {das}, // e.g. https://api.launchpad.net/1.0/ubuntu/xenial/armhf
		"metadata_override":  {metadata},
	}

	// Set up the URL
	u, err := url.Parse(apiURL)
	if err != nil {
		return err
	}
	u.Path = path.Join("livefs", "requestBuild")

	// Call the API
	resp, err := cli.oauthClient.Post(nil, &cli.credentials, u.String(), params)
	log.Println(err)
	log.Println(resp.Body)
	return err
}

// BuildMetadata returns the metadata override for the build
func (cli *Client) buildMetadata(boardID, osID string) (string, string, error) {
	// Get the metadata for the board and OS
	key := fmt.Sprintf("%s-%s", boardID, osID)
	meta := boards[key]

	// Serialize the data as JSON
	b, err := json.Marshal(meta)
	if err != nil {
		return "", "", err
	}
	return meta.DistroArchSeries, string(b), err
}
