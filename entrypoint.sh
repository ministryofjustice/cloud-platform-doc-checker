#!/bin/sh -l

go get -d -v ./...
go install -v ./...
go build .

./cloud-platform-doc-checker -team $1