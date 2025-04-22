main_src = ./cmd/service/main.go

ast:
	go tool compile -W $(main_src) > out_ast.txt

ssa:
	GOSSAFUNC=main go tool compile $(main_src)

asm:
	go tool compile -S $(main_src) > out_asm.txt

obj:
	go tool compile -o main.o $(main_src)
	go tool nm main.o > out_nm.txt

run:
	go run $(main_src)

clean:
	rm ssa.html main.o out*
