#!/bin/bash

# Любой пример go generate
go install github.com/golang/mock/mockgen@v1.6.0
go generate ./...
