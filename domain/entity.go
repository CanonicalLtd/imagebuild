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

// LiveFSBuild is the build asset on Launchpad
type LiveFSBuild struct {
	State       string   `json:"buildstate"`
	Link        string   `json:"web_link"`
	BuildLink   string   `json:"self_link"`
	LiveFS      string   `json:"livefs_link"`
	BuildLogURL string   `json:"build_log_url"`
	Archive     string   `json:"archive_link"`
	Duration    string   `json:"duration"`
	Started     string   `json:"date_started"`
	Built       string   `json:"datebuilt"`
	Title       string   `json:"title"`
	UniqueKey   string   `json:"unique_key"`
	Downloads   []string `json:"downloads"`
}

// BuildRequest is the request to initiate a build
type BuildRequest struct {
	BoardID  string   `json:"boardId"`
	OSID     string   `json:"osId"`
	Snaps    []string `json:"snaps"`
	Packages []string `json:"packages"`
}

// GetBuildRequest is the request to fetch a image build details
type GetBuildRequest struct {
	Link string `json:"link"`
}
