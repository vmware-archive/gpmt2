#!/bin/bash

SOURCE="github.com/pivotal-gss/gpmt2/cmd/gpmt"
TARGET="build/gpmt"

echo "building binary at build/gpmt"

go build -o "${TARGET}" "${SOURCE}"
