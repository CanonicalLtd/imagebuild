// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestWeb_StoreSearchHandler(t *testing.T) {
	settings, brdService := defaultsService()
	tests := []struct {
		name string
		url  string
		body string
		want int
	}{
		{"valid", "/v1/store/snaps/helloworld", `{"osId":"core16", "boardId":"raspberrypi3"}`, http.StatusOK},
		{"invalid-response", "/v1/store/snaps/invalid", `{"osId":"classic18.04", "boardId":"raspberrypi3"}`, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGET(`[{}]`)
			srv := NewWebService(settings, brdService)
			w := sendRequest("POST", tt.url, strings.NewReader(tt.body), srv)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}
		})
	}
}

func mockGET(body string) {
	// Mock the HTTP methods
	get = func(p string, h map[string]string) (*http.Response, error) {
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error get")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
}
