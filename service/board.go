// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package service

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/domain"
	"github.com/CanonicalLtd/imagebuild/launchpad"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Board interface for the service
type Board interface {
	List() []domain.Board
	Build(img *domain.BuildRequest) (string, error)
	GetLiveFSBuild(urlString string) (*domain.LiveFSBuild, error)
}

// BoardService implements the board service
type BoardService struct {
	boards    []domain.Board
	Settings  *config.Settings
	Launchpad launchpad.BuildClient
}

// NewBoardService creates a board service
func NewBoardService(settings *config.Settings, lp launchpad.BuildClient) *BoardService {
	brd := &BoardService{
		boards:    []domain.Board{},
		Settings:  settings,
		Launchpad: lp,
	}
	if err := brd.read(settings.BoardsPath); err != nil {
		log.Println("Error reading boards:", err)
	}
	return brd
}

// read reads the boards from the YAML file
func (brd *BoardService) read(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &brd.boards)
	return err
}

// List returns the list of boards
func (brd *BoardService) List() []domain.Board {
	return brd.boards
}

// Build starts an image build
func (brd *BoardService) Build(img *domain.BuildRequest) (string, error) {
	return brd.Launchpad.Build(img)
}

// GetLiveFSBuild retrieves the details of an image build
func (brd *BoardService) GetLiveFSBuild(urlString string) (*domain.LiveFSBuild, error) {
	return brd.Launchpad.GetLiveFSBuild(urlString)
}
