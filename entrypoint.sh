#!/bin/sh -l

go get -d -v ./...
go install -v ./...
go build .

./doc-checker -team $1