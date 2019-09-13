// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package main

import (
	"github.com/CanonicalLtd/imagebuild/config"
	"github.com/CanonicalLtd/imagebuild/service"
	"github.com/CanonicalLtd/imagebuild/web"
	"log"
)

func main() {
	settings := config.ParseArgs()
	brdSrv := service.NewBoardService(settings)
	srv := web.NewWebService(settings, brdSrv)
	log.Fatal(srv.Start())
}
