package launchpad

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/imagebuild/domain"
	"io"
	"net/http"
)

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

var doRequest = func(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}
