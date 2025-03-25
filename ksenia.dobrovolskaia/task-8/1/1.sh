#!/bin/bash

# Построение AST
go tool compile -W main.go

# Получение формы SSA + Генерация объектного файла
GOSSAFUNC=main go tool compile main.go > ssa.html

# Генерация исполняемого файла
go build
