// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package launchpad

import (
	"github.com/CanonicalLtd/imagebuild/domain"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var boards map[string]domain.Metadata

// readMetadata reads the build metadata for each board
func (cli *Client) readMetadata(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &boards)
	return err
}
