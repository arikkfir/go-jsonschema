#!/bin/sh

set -e
set -o errexit -o nounset

docker build --tag arikkfir/go-jsonschema:tools-latest --file Dockerfile --target tools .
