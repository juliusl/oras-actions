#!/bin/bash


# Scope the go path
GOPATH=$(echo $(go env | grep GOPATH) | awk -F 'GOPATH=' '{print $2}' | awk -F '"' '{print $2}')
ACTIONREPO=$(echo $PWD | awk -F "$GOPATH/" '{print $2}' | awk -F '/go/src/action' '{print $1}')
ACTIONREPO_GOPATH="$GOPATH/$ACTIONREPO/go"
GOPATH=ACTIONREPO_GOPATH
set -- 

if [-z $(go env) ]; then
    echo 'Install go to continue'
    exit 1
fi

# Install oras-go
go mod tidy