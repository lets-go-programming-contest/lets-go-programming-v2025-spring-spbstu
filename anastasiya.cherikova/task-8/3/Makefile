.PHONY: build build-custom clean

build:
	@echo "Building with default flags..."
	@go build -o main

build-custom:
	@echo "Building with custom ldflags..."
	@go build -ldflags "-X main.Version=1.0.0 \
		-X 'main.BuildTime=$(shell date +'%Y-%m-%d %H:%M:%S')' \
		-X main.CommitHash=$(shell git rev-parse --short HEAD)" \
		-o main

clean:
	@rm -f main

run: build
	@./main

run-custom: build-custom
	@./main