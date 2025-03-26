#!/usr/bin/env bash
set -x -e
main_path=$(dirname $0)/main.go
go generate
go run main.go
