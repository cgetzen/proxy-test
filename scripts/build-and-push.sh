#!/bin/sh
set -euo pipefail

SHA=$(git rev-parse --short HEAD)
docker buildx build . -t "cgetzen/proxy-test:${SHA}" --push
echo $SHA