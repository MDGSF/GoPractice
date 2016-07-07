#!/bin/bash

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install test

export GOPATH="$OLDGOPATH"

echo 'finished'
