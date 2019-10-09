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

func TestClient_buildMetadata(t *testing.T) {
	raspi2Core := `{"project":"ubuntu-core","subarch":"raspi2","image_format":"ubuntu-image"}`
	raspi3Core := `{"project":"ubuntu-core","subarch":"raspi3","image_format":"ubuntu-image"}`
	nucCore := `{"project":"ubuntu-core","image_format":"ubuntu-image"}`
	settings := &config.Settings{
		LPOwner:    "owner",
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
		BuildsPath: path.Join("..", config.DefaultBuildsPath),
	}

	doRequestValid := func(req *http.Request) (*http.Response, error) {
		return &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(`{}`))}, nil
	}

	type args struct {
		boardID string
		osID    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		{"valid-pi2-16", args{"raspberrypi2", "core16"}, "https://api.launchpad.net/1.0/ubuntu/xenial/armhf", "https://api.launchpad.net/devel/~owner/+livefs/ubuntu/xenial/ubuntu-core", raspi2Core, false},
		{"valid-pi2-16", args{"raspberrypi2", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/armhf", "https://api.launchpad.net/devel/~owner/+livefs/ubuntu/bionic/ubuntu-core", raspi2Core, false},
		{"valid-pi3-16", args{"raspberrypi3", "core16"}, "https://api.launchpad.net/1.0/ubuntu/xenial/armhf", "https://api.launchpad.net/devel/~owner/+livefs/ubuntu/xenial/ubuntu-core", raspi3Core, false},
		{"valid-pi3-18", args{"raspberrypi3", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/armhf", "https://api.launchpad.net/devel/~owner/+livefs/ubuntu/bionic/ubuntu-core", raspi3Core, false},
		{"valid-nuc-18", args{"intelnuc", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/amd64", "https://api.launchpad.net/devel/~owner/+livefs/ubuntu/bionic/ubuntu-core", nucCore, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doRequest = doRequestValid

			cli, _ := NewClient(settings, &MockAuthClient{})
			got, got1, got2, err := cli.buildMetadata(&domain.BuildRequest{BoardID: tt.args.boardID, OSID: tt.args.osID})
			if (err != nil) != tt.wantErr {
				t.Errorf("buildMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("buildMetadata() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("buildMetadata() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("buildMetadata() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestClient_requestBuild(t *testing.T) {
	settings := &config.Settings{
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
	}
	// Mock the HTTP request
	doRequest = mockDoRequest

	type fields struct {
		client OAuthClient
	}
	type args struct {
		das      string
		liveFS   string
		metadata string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid", fields{&MockAuthClient{}}, args{"https://api.launchpad.net/1.0/ubuntu/xenial/armhf", "+livefs/ubuntu/xenial/ubuntu-core", "{}"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, _ := NewClient(settings, tt.fields.client)
			if _, err := cli.requestBuild(tt.args.das, tt.args.liveFS, tt.args.metadata); (err != nil) != tt.wantErr {
				t.Errorf("requestBuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
