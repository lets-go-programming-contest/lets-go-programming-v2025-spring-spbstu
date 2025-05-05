lint:
	golangci-lint run

test:
	go test ./... -coverprofile cover.out
	go tool cover -o coverage.html -html=cover.out
