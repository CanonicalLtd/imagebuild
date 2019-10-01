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
	DefaultBuildsPath    = "builds.yaml"
	DefaultStoreURL      = "https://api.snapcraft.io/api/v1/"
	DefaultConsumer      = "image.build"
)

// Settings defines the application configuration
type Settings struct {
	Port          string
	DocRoot       string
	IndexTemplate string
	BoardsPath    string
	BuildsPath    string
	StoreURL      string
	LPOwner       string
	LPToken       string
	LPSecret      string
	LPConsumer    string
}

// ParseArgs checks the environment variables
func ParseArgs() *Settings {
	var (
		port = DefaultPort
	)

	if len(os.Getenv("LPTOKEN")) == 0 || len(os.Getenv("LPSECRET")) == 0 || len(os.Getenv("LPOWNER")) == 0 {
		log.Fatalln("The Launchpad access token, secret and owner must be supplied (LPTOKEN, LPSECRET, LPOWNER)")
	}

	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	return &Settings{
		Port:          port,
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
		BoardsPath:    DefaultBoardsPath,
		BuildsPath:    DefaultBuildsPath,
		StoreURL:      DefaultStoreURL,
		LPConsumer:    DefaultConsumer,
		LPToken:       os.Getenv("LPTOKEN"),
		LPSecret:      os.Getenv("LPSECRET"),
		LPOwner:       os.Getenv("LPOWNER"),
	}
}
