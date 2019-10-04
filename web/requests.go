// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"encoding/json"
	"github.com/CanonicalLtd/imagebuild/domain"
	"io"
	"log"
	"net/http"
)

func decodeBuildRequest(w http.ResponseWriter, r *http.Request) (*domain.BuildRequest, error) {
	defer r.Body.Close()

	// Decode the JSON body
	bld := domain.BuildRequest{}
	err := json.NewDecoder(r.Body).Decode(&bld)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("NoData", "No data supplied.", w)
		log.Println("Build request: no data supplied.")
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("BadData", err.Error(), w)
		log.Println("Build request:", err)
	}
	return &bld, err
}

func decodeGetBuildRequest(w http.ResponseWriter, r *http.Request) (*domain.GetBuildRequest, error) {
	defer r.Body.Close()

	// Decode the JSON body
	bld := domain.GetBuildRequest{}
	err := json.NewDecoder(r.Body).Decode(&bld)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("NoData", "No data supplied.", w)
		log.Println("GetBuild request: no data supplied.")
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("BadData", err.Error(), w)
		log.Println("GetBuild request:", err)
	}
	return &bld, err
}
