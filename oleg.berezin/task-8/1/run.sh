#!/bin/bash

go tool compile -W main.go > ast.txt

GOSSAFUNC=main go build main.go > ssa.html

go build
