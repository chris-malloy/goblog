#!/bin/bash
docker-compose -f docker-compose.yml -f env/local/docker-compose.local.yml down

# This script stops a local container build. This is the opposite of _up.sh