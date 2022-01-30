#!/bin/bash
usage() {
  echo "ERROR: Usage: $(basename $0) \$ROOTDIR: ";
  echo ROOTDIR is absolute path on your machine
 exit; }

ROOTDIR=${1:-"${PWD}"}
MOUNTPOINT=${2:-"/tree"} # This is fixed in the src code

if test -e ${ROOTDIR}; then
  if [[ $ROOTDIR == *"/../"* ]]; then
    usage
  fi
  docker build . \
    -t simple_server;
  docker run \
    --name peek-app -d -i \
    --rm \
    -p 8080:8000 \
    --mount type=bind,source="${PWD}/src",target=/app \
    --mount type=bind,source="${ROOTDIR}",target=${MOUNTPOINT} \
    simple_server
  set +x
else
  usage
fi
