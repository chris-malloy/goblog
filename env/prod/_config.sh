#!/bin/bash
docker-compose -f docker-compose.yml -f env/prod/docker-compose.prod.yml config

# This script runs the config command. Great for looking into environment variable issues.
# This script is non-destructive. It doesn't hurt to run it many times.