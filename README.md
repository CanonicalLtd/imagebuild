[![Build Status][travis-image]][travis-url]
[![Go Report Card][goreportcard-image]][goreportcard-url]
[![codecov][codecov-image]][codecov-url]
# Image Build

Ubuntu image building solution

 ## Build
 
 ### To test locally
 The project uses vendorized dependencies using `govendor`. Development has been done on minimum Go version 1.12.*.
 ```bash
 $ go get github.com/CanonicalLtd/imagebuild
 $ cd imagebuild
 $ GO111MODULES=on go build ./...
 ```
 
 ### With docker
 ```bash
 $ docker build -t imagebuild .
 $ docker run -p 8000:8000 imagebuild
 ```
 
 ## Run
 ```bash
 go run cmd/imagebuild/main.go
 ```

[travis-image]: https://travis-ci.org/CanonicalLtd/imagebuild.svg?branch=master
[travis-url]: https://travis-ci.org/CanonicalLtd/imagebuild
[goreportcard-image]: https://goreportcard.com/badge/github.com/CanonicalLtd/imagebuild
[goreportcard-url]: https://goreportcard.com/report/github.com/CanonicalLtd/imagebuild
[codecov-url]: https://codecov.io/gh/CanonicalLtd/imagebuild
[codecov-image]: https://codecov.io/gh/CanonicalLtd/imagebuild/branch/master/graph/badge.svg