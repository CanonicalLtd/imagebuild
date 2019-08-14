package main

import (
	"launchpad.net/ce-web/imagebuild/config"
	"launchpad.net/ce-web/imagebuild/web"
	"log"
)

func main() {
	settings := config.ParseArgs()
	srv := web.NewWebService(settings)
	log.Fatal(srv.Start())
}
