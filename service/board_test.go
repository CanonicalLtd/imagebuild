// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package service

import (
	"github.com/slimjim777/imagebuild/config"
	"github.com/slimjim777/imagebuild/domain"
	"path"
	"reflect"
	"testing"
)

func TestBoardService_List(t *testing.T) {
	brdOne := []domain.Board{{ID: "raspberrypi4", Name: "Raspberry Pi 4"}}
	brdNone := []domain.Board{}

	type fields struct {
		boards   []domain.Board
		Settings *config.Settings
	}
	tests := []struct {
		name   string
		fields fields
		want   []domain.Board
	}{
		{"valid-one", fields{brdOne, nil}, brdOne},
		{"valid-none", fields{brdNone, nil}, brdNone},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			brd := &BoardService{
				boards:   tt.fields.boards,
				Settings: tt.fields.Settings,
			}
			if got := brd.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoardService(t *testing.T) {
	settingsInvalid := &config.Settings{BoardsPath: "does.not.exist.yaml"}
	settingsValid := &config.Settings{BoardsPath: path.Join("..", config.DefaultBoardsPath)}

	type args struct {
		settings *config.Settings
	}
	tests := []struct {
		name  string
		args  args
		count int
	}{
		{"invalid-path", args{settingsInvalid}, 0},
		{"valid-path", args{settingsValid}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBoardService(tt.args.settings)
			if len(got.boards) != tt.count {
				t.Errorf("NewBoardService() = %v, want %v", len(got.boards), tt.count)
				return
			}
		})
	}
}
