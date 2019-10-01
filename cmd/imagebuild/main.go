// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package main

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/launchpad"
	"github.com/CanonicalLtd/imagebuild/service"
	"github.com/CanonicalLtd/imagebuild/web"
	"log"
)

func main() {
	// Parse the command-line arguments
	settings := config.ParseArgs()

	// Set up the dependency chain
	authClient := launchpad.NewOAuthClient(settings)
	lp, err := launchpad.NewClient(settings, authClient)
	if err != nil {
		log.Fatalln("Error initialising Launchpad client:", err)
	}
	brdSrv := service.NewBoardService(settings, lp)
	srv := web.NewWebService(settings, brdSrv)

	// Start the web service
	log.Fatal(srv.Start())
}
