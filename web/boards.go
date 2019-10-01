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

// BoardsList list the available boards for enablement
func (srv Web) BoardsList(w http.ResponseWriter, r *http.Request) {
	boards := srv.BoardSrv.List()
	formatBoardsResponse(boards, w)
}

// Build starts a image build
func (srv Web) Build(w http.ResponseWriter, r *http.Request) {
	bld, err := decodeBuildRequest(w, r)
	if err != nil {
		return
	}

	buildURL, err := srv.BoardSrv.Build(bld)
	if err != nil {
		formatStandardResponse("Build", err.Error(), w)
	}
	formatBuildResponse(buildURL, w)
}

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
