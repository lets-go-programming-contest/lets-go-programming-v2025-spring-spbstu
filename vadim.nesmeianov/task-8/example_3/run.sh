#!/bin/bash

go build -ldflags="-X 'main.Version=v1.0.1'"
./part3