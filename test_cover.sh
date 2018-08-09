#!/usr/bin/env bash

set -e

for d in $(go list ./... | grep -v vendor); do
    go test -coverprofile=profile.out -covermode=atomic -v $d
    if [ -f profile.out ]; then
        rm profile.out
    fi
done