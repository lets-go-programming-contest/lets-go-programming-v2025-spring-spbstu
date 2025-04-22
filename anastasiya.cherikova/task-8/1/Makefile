.PHONY: all ast ssa asm obj exe clean

all: ast ssa asm obj exe

ast:
	@echo "Generating AST..."
	@GO111MODULE=off go tool compile -W main.go > ast.txt

ssa:
	@echo "Generating SSA..."
	@GOSSAFUNC=main go tool compile main.go 2>&1 > ssa.html

asm:
	@echo "Generating assembly..."
	@go build -gcflags="-S" main.go 2>&1

obj:
	@echo "Generating object file..."
	@go tool compile -o main.o main.go

exe:
	@echo "Building executable..."
	@go build -o main

clean:
	@rm -f main main.o main.s ssa.html ast.txt

run: exe
	@./main