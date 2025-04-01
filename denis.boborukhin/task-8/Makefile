main_src = ./cmd/service/main.go
ast_example_src = ./cmd/ast_example/main.go

build-ast:
	go build $(ast_example_src)

get-ast:
	go tool compile -W $(ast_example_src)

get-asm:
	go build -gcflags="-S" $(ast_example_src)

get-ssa:
	GOSSAFUNC=main go tool compile $(ast_example_src) > ssa.html

tags:
	go run -tags tags $(main_src)

ld:
	go run -ldflags="-X 'main.Version=ldflags'" $(main_src)

generate:
	go generate $(main_src)

build:
	go run $(main_src)

clean:
	rm ssa.html main main.o ./internal/data/hello.txt ./internal/data/world.txt
