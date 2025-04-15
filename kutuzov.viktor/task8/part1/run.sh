#!/bin/bash

# Построение AST
go tool compile -W main.go > ast.txt

# Получение ассемблера + исполняемый файл
go build -gcflags="-S" main.go 2> asm.txt # Компилятор Go выводит техническую информацию (включая ассемблерный код) в stderr, чтобы не смешивать её с обычным выводом программы.

# Получение формы SSA 
GOSSAFUNC=main go tool compile main.go > ssa.html
firefox ssa.html