#!/bin/bash

# Добавить в main.go тэги и проверить работоспособность
go build -o simple
go build -tags pro -o tagged
