package main

import (
	"github.com/slimjim777/imagebuild/config"
	"github.com/slimjim777/imagebuild/web"
	"log"
)

func main() {
	settings := config.ParseArgs()
	srv := web.NewWebService(settings)
	log.Fatal(srv.Start())
}
