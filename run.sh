#!/bin/sh

cd "$(dirname "$0")/src" || exit 1

go run . "$@"