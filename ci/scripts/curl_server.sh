#!/bin/sh
# Curls an endpoint for a status code. If 400 or 500 codes are returned, should exit with process 1.

if [ "$1" = "" ]; then
    echo "Unable to parse load balancer DNS."
    exit 1
fi

if [ "$2" = "" ]; then
    echo "Unable to parse status route."
    exit 1
fi

export LB_DNS=$1
export STATUS_ROUTE=$2

# matchers will be compared against the status code
export CLIENT_ERROR_MATCHER=4
export SERVER_ERROR_MATCHER=5

# curl the server to see if it's up
# -k to ignore ssl handshake
# -L to allow forwarding (such as http -> https)
# -s for quiet output so we only get what's specified by -w
# -w specify what to output
# -o set the output file. in this case we want to toss out everything except for what we pass to -w

STATUS_CODE=$(curl -o /dev/null -kLsw "%{http_code}" ${LB_DNS}${STATUS_ROUTE})

if echo "${STATUS_CODE:0:1}" | grep "${CLIENT_ERROR_MATCHER}" >/dev/null; then
    echo "status code reflects client error:" "${STATUS_CODE}"
    exit 1
elif echo "${STATUS_CODE:0:1}" | grep "${SERVER_ERROR_MATCHER}" >/dev/null; then
    echo "status code reflects server error:" "${STATUS_CODE}"
    exit 1
else
    echo "status code reflects the app is up and running:" "${STATUS_CODE}"
    exit 0
fi
