#!/bin/bash

# Воспроизвести пример применения ldflags
go build -o development
go build -ldflags="-X 'main.Version=v1.0.0'" -o v1

