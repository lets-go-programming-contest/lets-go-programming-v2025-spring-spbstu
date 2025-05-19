#!/bin/bash

go build -o simple
go build -ldflags="-s -w" -o ldflagged #stripping
