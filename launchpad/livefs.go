// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/imagebuild/domain"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
	if liveBuild.State == "Successfully built" {
		err = cli.getBuildURL(u, liveBuild)
	}
	return liveBuild, err
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

func decodeLiveFSBuild(r *http.Response) (*domain.LiveFSBuild, error) {
	defer r.Body.Close()

	// Decode the JSON body
	bld := domain.LiveFSBuild{}
	err := json.NewDecoder(r.Body).Decode(&bld)
	switch {
	// Check we have some data
	case err == io.EOF:
		return nil, fmt.Errorf("livefs build request: no data supplied")
		// Check for parsing errors
	case err != nil:
		return nil, fmt.Errorf("livefs build request: %v", err)
	}
	return &bld, err
}

func decodeLiveFSAssets(r *http.Response) ([]string, error) {
	var assets []string
	defer r.Body.Close()

	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&assets)
	switch {
	// Check we have some data
	case err == io.EOF:
		return nil, fmt.Errorf("livefs assets request: no data supplied")
		// Check for parsing errors
	case err != nil:
		return nil, fmt.Errorf("livefs assets request: %v", err)
	}
	return assets, err
}
