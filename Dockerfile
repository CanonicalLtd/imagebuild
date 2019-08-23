FROM golang:1.12 as builder1
COPY . ./src/github.com/slimjim777/imagebuild
WORKDIR /go/src/github.com/slimjim777/imagebuild
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/imagebuild cmd/imagebuild/main.go

FROM node:8-alpine as builder2
COPY webapp .
WORKDIR /
RUN npm install
RUN npm rebuild node-sass
RUN npm run build

# Copy the built applications to the docker image
FROM ubuntu:18.04
WORKDIR /root/
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=builder1 /go/bin/imagebuild .
COPY --from=builder2 build/ ./static/
COPY boards.yaml .
EXPOSE 8000
ENTRYPOINT ./imagebuild