// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package config

import (
	"os"
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	wantDefault := &Settings{
		Port:          DefaultPort,
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
		BoardsPath:    DefaultBoardsPath,
		StoreURL:      DefaultStoreURL,
	}
	want9000 := &Settings{
		Port:          ":9000",
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
		BoardsPath:    DefaultBoardsPath,
		StoreURL:      DefaultStoreURL,
	}

	tests := []struct {
		name string
		want *Settings
	}{
		{"valid", wantDefault},
		{"valid-9000", want9000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.Port != DefaultPort {
				os.Setenv("PORT", tt.want.Port)
			}

			if got := ParseArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
