#!/bin/sh
export GOPATH=$(pwd)
export GOBIN=$GOPATH/bin

#ginkgo watch -r -notify src/goblog.com/_unittests

# run this command if you prefer to not run the suite with ginkgo
 go test -v goblog.com/_unittests
