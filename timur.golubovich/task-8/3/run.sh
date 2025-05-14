#!/bin/bash

# Воспроизвести пример применения ldflags
go build -o simple
go build -ldflags="-s -w" -o ldflagged #stripping
