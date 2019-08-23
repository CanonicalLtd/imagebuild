# Image Build

Prototype for the image build solution

 ## Build
 
 ### To test locally
 The project uses vendorized dependencies using `govendor`. Development has been done on minimum Go version 1.12.*.
 ```bash
 $ go get github.com/slimjim777/imagebuild
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