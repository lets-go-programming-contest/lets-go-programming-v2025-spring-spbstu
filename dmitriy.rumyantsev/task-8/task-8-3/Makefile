main_src = ./cmd/service

export VERSION := v1.0.0

ldflags:
	go run -ldflags="-X 'main.Version=$(VERSION)'" $(main_src)
