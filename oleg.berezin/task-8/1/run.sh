#!/bin/bash

# Построение AST
go tool compile -W main.go > ast.txt

# Получение формы SSA + Генерация объектного файла
GOSSAFUNC=main go build main.go > ssa.html

# Генерация исполняемого файла
go build
