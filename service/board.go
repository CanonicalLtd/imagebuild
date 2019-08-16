package service

import (
	"github.com/slimjim777/imagebuild/config"
	"github.com/slimjim777/imagebuild/domain"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Board interface for the service
type Board interface {
	List() []domain.Board
}

// BoardService implements the board service
type BoardService struct {
	boards   []domain.Board
	Settings *config.Settings
}

// NewBoardService creates a board service
func NewBoardService(settings *config.Settings) *BoardService {
	brd := &BoardService{
		boards:   []domain.Board{},
		Settings: settings,
	}
	if err := brd.read(settings.BoardsPath); err != nil {
		log.Println("Error reading boards:", err)
	}
	return brd
}

// Read reads the boards from the YAML file
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
	if brd.boards == nil {
		return []domain.Board{}
	}
	return brd.boards
}
