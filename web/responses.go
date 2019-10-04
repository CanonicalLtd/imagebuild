// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"encoding/json"
	"github.com/CanonicalLtd/imagebuild/domain"
	"log"
	"net/http"
)

// JSONHeader is the header for JSON responses
const JSONHeader = "application/json; charset=UTF-8"

// StandardResponse is the JSON response from an API method, indicating success or failure.
type StandardResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// BoardsResponse is the JSON response to list snaps
type BoardsResponse struct {
	StandardResponse
	Boards []domain.Board `json:"boards"`
}

// BuildResponse is the JSON response to list snaps
type BuildResponse struct {
	StandardResponse
	BuildURL string `json:"buildURL"`
}

// LiveFSBuildResponse is the JSON response to list snaps
type LiveFSBuildResponse struct {
	StandardResponse
	Build domain.LiveFSBuild `json:"build"`
}

// formatStandardResponse returns a JSON response from an API method, indicating success or failure
func formatStandardResponse(code, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := StandardResponse{Code: code, Message: message}

	if len(code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatBoardsResponse returns a JSON response from a snap list API method
func formatBoardsResponse(boards []domain.Board, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := BoardsResponse{StandardResponse{}, boards}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatBuildResponse returns a JSON response from a snap list API method
func formatBuildResponse(b string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := BuildResponse{StandardResponse{}, b}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatLiveFSBuildResponse returns a JSON response from a snap list API method
func formatLiveFSBuildResponse(b *domain.LiveFSBuild, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := LiveFSBuildResponse{StandardResponse{}, *b}

	// Encode the response as JSON
	encodeResponse(w, response)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the response:", err)
	}
}
