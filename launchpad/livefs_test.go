// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/domain"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestClient_GetLiveFSBuild(t *testing.T) {
	settings := &config.Settings{
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
	}
	type fields struct {
		oauthClient OAuthClient
	}
	// Mock the HTTP request
	doRequestValid := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`{"buildstate": "Currently building"}`))}, nil
	}
	doRequestEmpty := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(``))}, nil
	}
	doRequestBad := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`\u1000`))}, nil
	}

	resp1 := &domain.LiveFSBuild{State: "Currently building"}

	tests := []struct {
		name    string
		fields  fields
		rawURL  string
		request func(req *http.Request) (*http.Response, error)
		want    *domain.LiveFSBuild
		wantErr bool
	}{
		{"valid", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestValid, resp1, false},
		{"empty-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestEmpty, nil, true},
		{"bad-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestBad, nil, true},
		{"error-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", mockDoRequestError, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doRequest = tt.request
			cli, _ := NewClient(settings, tt.fields.oauthClient)
			got, err := cli.GetLiveFSBuild(tt.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLiveFSBuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLiveFSBuild() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_getBuildURL(t *testing.T) {
	settings := &config.Settings{
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
	}
	// Mock the HTTP request
	doRequestValid := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`["abc", "def"]`))}, nil
	}
	doRequestEmpty := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(``))}, nil
	}
	doRequestBad := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`\u1000`))}, nil
	}

	resp1 := []string{"abc", "def"}

	type fields struct {
		oauthClient OAuthClient
	}

	tests := []struct {
		name    string
		fields  fields
		rawURL  string
		request func(req *http.Request) (*http.Response, error)
		want    []string
		wantErr bool
	}{
		{"valid", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestValid, resp1, false},
		{"empty-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestEmpty, nil, true},
		{"bad-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", doRequestBad, nil, true},
		{"error-response", fields{&MockAuthClient{}}, "https://api.launchpad.net/devel/~jamesj/+livefs/ubuntu/xenial/ubuntu-core/+build/184825", mockDoRequestError, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doRequest = tt.request
			cli, _ := NewClient(settings, tt.fields.oauthClient)
			u, _ := url.Parse(tt.rawURL)

			liveBuild := &domain.LiveFSBuild{}

			err := cli.getBuildURL(u, liveBuild)
			if (err != nil) != tt.wantErr {
				t.Errorf("getBuildURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(liveBuild.Downloads, tt.want) {
				t.Errorf("getBuildURL() got = %v, want %v", liveBuild.Downloads, tt.want)
			}
		})
	}
}
