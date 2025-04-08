main_src = ./cmd/service

generate:
	go generate ./...

run:
	go run $(main_src)

all:
	go generate ./...
	go run $(main_src)

clean:
	rm -r internal/mocks
