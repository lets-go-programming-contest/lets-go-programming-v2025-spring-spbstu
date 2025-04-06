main_src = ./cmd/service

generate:
	go generate $(main_src)

run:
	go run $(main_src)

all:
	go generate $(main_src)
	go run $(main_src)

clean:
	rm -r internal/mocks
