#!/bin/bash
docker-compose -f docker-compose.yml -f env/prod/docker-compose.prod.yml up --build

# This script stops a local container build. This is the opposite of _up.sh