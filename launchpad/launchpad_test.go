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
