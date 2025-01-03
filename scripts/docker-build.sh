#!/bin/sh

set -e
set -o errexit -o nounset

docker build --tag arikkfir/go-jsonschema:latest --file Dockerfile --target go-jsonschema .
