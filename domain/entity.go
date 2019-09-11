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
