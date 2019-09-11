// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

package main

import (
	"github.com/slimjim777/imagebuild/config"
	"github.com/slimjim777/imagebuild/service"
	"github.com/slimjim777/imagebuild/web"
	"log"
)

func main() {
	settings := config.ParseArgs()
	brdSrv := service.NewBoardService(settings)
	srv := web.NewWebService(settings, brdSrv)
	log.Fatal(srv.Start())
}
