// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package config

import (
	"log"
	"os"
)

// Default settings
const (
	DefaultPort          = ":8000"
	DefaultDocRoot       = "static"
	DefaultIndexTemplate = "index.html"
	DefaultBoardsPath    = "boards.yaml"
	DefaultStoreURL      = "https://api.snapcraft.io/api/v1/"
)

// Settings defines the application configuration
type Settings struct {
	Port          string
	DocRoot       string
	IndexTemplate string
	BoardsPath    string
	StoreURL      string
	LPToken       string
	LPSecret      string
}

// ParseArgs checks the environment variables
func ParseArgs() *Settings {
	var (
		port = DefaultPort
	)

	if len(os.Getenv("LPTOKEN")) == 0 || len(os.Getenv("LPSECRET")) == 0 {
		log.Fatalln("The Launchpad access token and secret must be supplied (LPTOKEN, LPSECRET)")
	}

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	return &Settings{
		Port:          port,
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
		BoardsPath:    DefaultBoardsPath,
		StoreURL:      DefaultStoreURL,
		LPToken:       os.Getenv("LPTOKEN"),
		LPSecret:      os.Getenv("LPSECRET"),
	}
}
