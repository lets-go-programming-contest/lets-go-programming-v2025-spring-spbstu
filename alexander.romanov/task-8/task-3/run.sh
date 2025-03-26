#!/usr/bin/env bash
set -x -e
main_path=$(dirname $0)/main.go
go build -ldflags="-X main.Version=1.0.0" -o idflags_executable $main_path
echo "Executable with idflags generated as idflags_executable"
