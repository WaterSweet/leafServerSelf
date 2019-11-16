#!/bin/bash

curDir=$(pwd)
export GOPATH=$GOPATH:$curDir/src/common
export GOPATH=$GOPATH:$curDir/src/server

go build -o $curDir/main.go

echo "done"
