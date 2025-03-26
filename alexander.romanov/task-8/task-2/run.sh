#!/usr/bin/env bash
set -x
main_path=$(dirname $0)/main.go
echo "COMPILING WITH TAGS:"
go build -tags mytag -o with
./with
echo "COMPILING WITHOUT TAGS:"
go build -o with.out
./with.out
