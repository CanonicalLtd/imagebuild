// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/gomodule/oauth1/oauth"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const apiURL = "https://api.launchpad.net/devel"

// BuildClient defines the interface for the Launchpad client for building images
type BuildClient interface {
	Build(boardID, osID string) (string, error)
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
func (cli *Client) Build(boardID, osID string) (string, error) {
	// Get the metadata for the selection
	distroArchSeries, liveFS, meta, err := cli.buildMetadata(boardID, osID)
	if err != nil {
		return "", err
	}

	// Start the build
	return cli.requestBuild(distroArchSeries, liveFS, meta)
}

// requestBuild starts a new build as per https://launchpad.net/+apidoc/devel.html#livefs-requestBuild
func (cli *Client) requestBuild(das, liveFS, metadata string) (string, error) {
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
		return "", err
	}
	u.Path = path.Join(u.Path, "~"+cli.settings.LPOwner, liveFS)

	// Call the API
	resp, err := cli.httpDo("POST", u, params)
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error submitting %d build: %v", resp.StatusCode, string(bodyBytes))
	}

	return resp.Header.Get("Location"), err
}

// httpDo performs the HTTP method, setting the OAuth authorization header
func (cli *Client) httpDo(method string, u *url.URL, form url.Values) (*http.Response, error) {
	// Create the request
	req, err := http.NewRequest(method, u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	// Set the request header with the correct credentials, and the correct consumer
	if err := cli.oauthClient.SetAuthorizationHeader(req.Header, &cli.credentials, method, u, form); err != nil {
		return nil, err
	}

	return doRequest(req)
}

// BuildMetadata returns the metadata override for the build
func (cli *Client) buildMetadata(boardID, osID string) (string, string, string, error) {
	// Get the metadata for the board and OS
	key := fmt.Sprintf("%s-%s", boardID, osID)
	meta := boards[key]

	// Serialize the data as JSON
	b, err := json.Marshal(meta)
	if err != nil {
		return "", "", "", err
	}
	return meta.DistroArchSeries, meta.LiveFS, string(b), err
}

var doRequest = func(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}
