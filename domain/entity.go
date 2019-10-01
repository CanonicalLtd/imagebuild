// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package domain

// Board defines a board for enablement
type Board struct {
	ID   string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
	OS   []struct {
		ID      string `yaml:"id" json:"id"`
		Type    string `yaml:"type" json:"type"`
		Version string `yaml:"version" json:"version"`
	} `yaml:"os" json:"os"`
}

// Metadata describes the parameters for the API for a board and OS
type Metadata struct {
	DistroArchSeries string `yaml:"distro_arch_series" json:"-"`
	LiveFS           string `yaml:"livefs" json:"-"`
	Project          string `yaml:"project" json:"project,omitempty"`
	SubArch          string `yaml:"subarch" json:"subarch,omitempty"`
	ImageFormat      string `yaml:"image_format" json:"image_format,omitempty"`
}

// BuildRequest is the request to initiate a build
type BuildRequest struct {
	BoardID  string   `json:"boardId"`
	OSID     string   `json:"osId"`
	Snaps    []string `json:"snaps"`
	Packages []string `json:"packages"`
}
