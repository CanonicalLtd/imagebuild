// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/domain"
	"path"
	"testing"
)

func TestClient_Build(t *testing.T) {
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
		boardID string
		osID    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"valid", fields{&MockAuthClient{}}, args{"raspberrypi2", "core16"}, false},
		{"invalid", fields{&MockAuthClient{}}, args{"raspberrypi2", "core16"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				doRequest = mockDoRequestError
			}

			cli, _ := NewClient(settings, tt.fields.client)
			if _, err := cli.Build(&domain.BuildRequest{BoardID: tt.args.boardID, OSID: tt.args.osID}); (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_buildMetadata(t *testing.T) {
	raspi2Core := `{"project":"ubuntu-core","subarch":"raspi2","image_format":"ubuntu-image"}`
	raspi3Core := `{"project":"ubuntu-core","subarch":"raspi3","image_format":"ubuntu-image"}`
	nucCore := `{"project":"ubuntu-core","image_format":"ubuntu-image"}`
	settings := &config.Settings{
		DocRoot:    "../static",
		BoardsPath: path.Join("..", config.DefaultBoardsPath),
		BuildsPath: path.Join("..", config.DefaultBuildsPath),
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
		{"valid-pi2-16", args{"raspberrypi2", "core16"}, "https://api.launchpad.net/1.0/ubuntu/xenial/armhf", "+livefs/ubuntu/xenial/ubuntu-core", raspi2Core, false},
		{"valid-pi2-16", args{"raspberrypi2", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/armhf", "+livefs/ubuntu/bionic/ubuntu-core", raspi2Core, false},
		{"valid-pi3-16", args{"raspberrypi3", "core16"}, "https://api.launchpad.net/1.0/ubuntu/xenial/armhf", "+livefs/ubuntu/xenial/ubuntu-core", raspi3Core, false},
		{"valid-pi3-18", args{"raspberrypi3", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/armhf", "+livefs/ubuntu/bionic/ubuntu-core", raspi3Core, false},
		{"valid-nuc-18", args{"intelnuc", "core18"}, "https://api.launchpad.net/1.0/ubuntu/bionic/amd64", "+livefs/ubuntu/bionic/ubuntu-core", nucCore, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
