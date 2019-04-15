#!/bin/bash

## clean
rm -rf build/

## for Mac
(
    GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w" -v -o build/darwin-amd64/pwgen cmd/pw/*.go
    cd ./build/darwin-amd64 || exit
    tar cfz pw-darwin-amd64.tar.gz pwgen
)

## for windows
(
    GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w" -v -o build/windows-amd64/pwgen.exe cmd/pw/*.go
    cd ./build/windows-amd64 || exit
    zip -q pw-windows-amd64.tar.gz pwgen.exe
)

## for linux
(
    GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-s -w" -v -o build/linux-amd64/pwgen cmd/pw/*.go
    cd ./build/linux-amd64 || exit
    tar cfz pw-linux-amd64.tar.gz pwgen
)
