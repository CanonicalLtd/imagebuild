// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/imagebuild/domain"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Launchpad constants
const (
	SuccessfullyBuilt = "Successfully built"
)

// getOrCreateLiveFS makes sure that the liveFS object exists in Launchpad
func (cli *Client) getOrCreateLiveFS(rawURL, name, distroSeries, ownerLink string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	// Fetch the livefs object
	if err := cli.getLiveFS(u); err == nil {
		// Object exists
		return nil
	}

	// Create the livefs record
	return cli.createLiveFS(name, distroSeries, ownerLink)
}

// getLiveFS fetches a liveFS object exists in Launchpad
func (cli *Client) getLiveFS(u *url.URL) error {
	// Call the API
	resp, err := cli.httpDo("GET", u, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status %d when fetching livefs %s", resp.StatusCode, u.String())
	}
	return nil
}

// createLiveFS creates a liveFS object exists in Launchpad
func (cli *Client) createLiveFS(name, distroSeries, ownerLink string) error {
	// Set the parameters
	params := map[string][]string{
		"ws.op":         {"new"},
		"distro_series": {distroSeries}, // e.g. https://api.launchpad.net/1.0/ubuntu/xenial
		"metadata":      {"{}"},
		"name":          {name},
		"owner":         {ownerLink},
	}

	u, err := url.Parse(apiURL + "/livefses")
	if err != nil {
		return err
	}

	// Call the API
	resp, err := cli.httpDo("POST", u, params)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		return fmt.Errorf("status %d when creating livefs %s", resp.StatusCode, u.String())
	}
	return nil
}

// getBuildURL fetches the download URL if the build is completed successfully
func (cli *Client) getBuildURL(u *url.URL, liveBuild *domain.LiveFSBuild) error {
	// Set the query for the assets
	q := u.Query()
	q.Add("ws.op", "getFileUrls")
	u.RawQuery = q.Encode()

	// Get the assets from the API
	resp, err := cli.httpDo("GET", u, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		return fmt.Errorf("error fetching %d build assets: %v", resp.StatusCode, string(bodyBytes))
	}

	// Parse the response
	assets, err := decodeLiveFSAssets(resp)
	if err != nil {
		return err
	}
	liveBuild.Downloads = assets
	return nil
}

// BuildMetadata returns the metadata override for the build
func (cli *Client) buildMetadata(img *domain.BuildRequest) (string, string, string, error) {
	// Get the metadata for the board and OS
	meta := cli.BoardMeta(img.BoardID, img.OSID)

	// Set the snaps and packages, if they are provided
	if img.Snaps != nil && len(img.Snaps) > 0 {
		meta.Snaps = img.Snaps
	}
	if img.Packages != nil && len(img.Packages) > 0 {
		meta.Packages = img.Packages
	}

	// Generate the distro_arch_series
	das := meta.DistroSeries + "/" + meta.Arch

	// Generate the owner link
	ownerLink := apiURL + "/~" + cli.settings.LPOwner

	// Generate the livefs URL
	liveFSURL := apiURL + "/~" + cli.settings.LPOwner + "/" + meta.LiveFS

	// Make sure that we create the livefs record, if needed
	if err := cli.getOrCreateLiveFS(liveFSURL, meta.Project, meta.DistroSeries, ownerLink); err != nil {
		return "", "", "", err
	}

	// Serialize the data as JSON
	b, err := json.Marshal(meta)
	if err != nil {
		return "", "", "", err
	}
	return das, liveFSURL, string(b), err
}

// requestBuild starts a new build as per https://launchpad.net/+apidoc/devel.html#livefs-requestBuild
func (cli *Client) requestBuild(das, liveFS, metadata string) (string, error) {
	// Set the parameters
	params := map[string][]string{
		"ws.op":              {"requestBuild"},
		"pocket":             {"Updates"},
		"archive":            {"https://api.launchpad.net/1.0/ubuntu/+archive/primary"},
		"distro_arch_series": {das}, // e.g. https://api.launchpad.net/1.0/ubuntu/xenial/armhf
		"metadata_override":  {metadata},
	}

	// Set up the URL
	u, err := url.Parse(liveFS)
	if err != nil {
		return "", err
	}

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
