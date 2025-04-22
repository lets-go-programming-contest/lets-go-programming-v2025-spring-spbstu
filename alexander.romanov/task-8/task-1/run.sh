#!/usr/bin/env bash
set -x -e
main_path=$(dirname $0)/main.go
file=$(dirname $0)/example.txt
gofile=$(dirname $0)/example.go
echo "AST GENERATION:"
go run $main_path $file
echo "SSA GENERATION:"
GOSSAFUNC=main go build -o $(mktemp /tmp/exampl-ssa.tmpXXX) ./ast/example.go
echo "ASSEMBLY GENERATION:"
go build -gcflags="-S" -o $(mktemp /tmp/exampl-asm.tmpXXX) ./ast/example.go
echo "OBJECT GENERATION:"
go tool compile $gofile
echo "GENERATED $(basename $gofile .go).o"
echo "EXECUTABLE GENERATION:"
go build $main_path
