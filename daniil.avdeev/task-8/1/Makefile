fib_src = ./cmd/service/main.go

all:
	go build $(fib_src) 
	go tool compile -W $(fib_src) > fib_ast
	go build -gcflags="-S" $(fib_src) > fib.s
	GOSSAFUNC=fib go tool compile $(fib_src) > fub_ssa.html
