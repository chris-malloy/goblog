#!/bin/bash

if [[ "$#" -ne 1 ]]; then
    echo "Unable to parse environment."
    exit -1
fi

# Build containers
docker-compose -f docker-compose.yml -f env/$1/docker-compose.$1.yml build

# This is the least invasive way to execute our integration tests. We do not want to
# put anything in the container that smells like test scripts. The sleep below waits on
# the entire docker suite to load before running the server--which requires a database
# to be live before it can successfully wake up
docker-compose -f docker-compose.yml \
-f env/$1/docker-compose.$1.yml run \
--entrypoint "/bin/bash -c \"sleep 3; /go/bin/server & go test -v goblog.com/_inttests \"" goblog

# cleanup
docker-compose -f docker-compose.yml -f env/$1/docker-compose.$1.yml down
