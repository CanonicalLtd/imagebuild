// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"github.com/slimjim777/imagebuild/config"
	"net/http"
	"strings"
	"testing"
)

func TestWeb_Index(t *testing.T) {
	settings, brdService := defaultsService()

	tests := []struct {
		name     string
		filePath string
		want     int
	}{
		{"valid", config.DefaultIndexTemplate, http.StatusOK},
		{"invalid-template", "does-not-exist.html", http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings.IndexTemplate = tt.filePath
			srv := NewWebService(settings, brdService)

			w := sendRequest("GET", "/", strings.NewReader(""), srv)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}
		})
	}
}
