// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"net/http"
	"strings"
	"testing"
)

func TestWeb_BoardsList(t *testing.T) {
	settings, brdService := defaultsService()

	tests := []struct {
		name    string
		want    int
		wantErr string
	}{
		{"valid", http.StatusOK, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(settings, brdService)

			w := sendRequest("GET", "/v1/boards", strings.NewReader(""), srv)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseBoardsResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.BoardsList() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestWeb_Build(t *testing.T) {
	settings, brdService := defaultsService()

	d1 := `{"boardId": "raspberrypi2", "osId": "core16"}`
	d2 := `\u1000`
	d3 := `{"boardId": "error", "osId": "core16"}`

	tests := []struct {
		name    string
		data    string
		want    int
		wantErr string
	}{
		{"valid", d1, http.StatusOK, ""},
		{"invalid-data", d2, http.StatusBadRequest, "BadData"},
		{"invalid-data", "", http.StatusBadRequest, "NoData"},
		{"error-build", d3, http.StatusBadRequest, "Build"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(settings, brdService)

			w := sendRequest("POST", "/v1/build", strings.NewReader(tt.data), srv)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseBuildResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.Build() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}
