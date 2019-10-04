// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
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
		return
	}
	formatBuildResponse(buildURL, w)
}

// GetLiveFSBuild fetches a liveFS build details
func (srv Web) GetLiveFSBuild(w http.ResponseWriter, r *http.Request) {
	bld, err := decodeGetBuildRequest(w, r)
	if err != nil {
		return
	}

	build, err := srv.BoardSrv.GetLiveFSBuild(bld.Link)
	if err != nil {
		formatStandardResponse("Build", err.Error(), w)
		return
	}
	formatLiveFSBuildResponse(build, w)
}
