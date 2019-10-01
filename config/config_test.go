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
		LPConsumer:    DefaultConsumer,
		LPOwner:       "johndoe",
	}
	want9000 := &Settings{
		Port:          ":9000",
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
		BoardsPath:    DefaultBoardsPath,
		StoreURL:      DefaultStoreURL,
		LPConsumer:    DefaultConsumer,
		LPOwner:       "johndoe",
	}

	tests := []struct {
		name     string
		want     *Settings
		lpToken  string
		lpSecret string
		lpOwner  string
	}{
		{"valid", wantDefault, "abc", "def", "johndoe"},
		{"valid-9000", want9000, "abc", "def", "johndoe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.lpToken) > 0 {
				os.Setenv("LPTOKEN", tt.lpToken)
				tt.want.LPToken = tt.lpToken
			}
			if len(tt.lpSecret) > 0 {
				os.Setenv("LPSECRET", tt.lpSecret)
				tt.want.LPSecret = tt.lpSecret
			}
			if len(tt.lpOwner) > 0 {
				os.Setenv("LPOWNER", tt.lpOwner)
				tt.want.LPOwner = tt.lpOwner
			}
			if tt.want.Port != DefaultPort {
				os.Setenv("PORT", tt.want.Port)
			}

			if got := ParseArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
