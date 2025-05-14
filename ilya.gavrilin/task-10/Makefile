.PHONY: proto build run-grpc run-rest run-client format-proto

install-tools:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
# cheap exchange for buf lint
format-proto:
	find ./api/proto -name "*.proto" -exec clang-format -i {} \;

proto:
	mkdir -p gen/proto
	protoc -I. \
		-I$(shell go env GOPATH)/src/github.com/googleapis/googleapis \
		-I$(shell go env GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=gen/proto --go_opt=paths=source_relative \
		--go-grpc_out=gen/proto --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=gen/proto --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=api/swagger \
		--validate_out="lang=go,paths=source_relative:gen/proto" \
		api/proto/contact_manager/v1/contact_manager.proto
	# Generate documentation
	protoc -I. \
		-I$(shell go env GOPATH)/src/github.com/googleapis/googleapis \
		-I$(shell go env GOPATH)/src/github.com/envoyproxy/protoc-gen-validate \
		--doc_out=docs/gen --doc_opt=markdown,api.md \
		api/proto/contact_manager/v1/contact_manager.proto

# Build everything
build:
	go build -o bin/grpc_server ./cmd/grpc_server
	go build -o bin/rest_server ./cmd/rest_server
	go build -o bin/grpc_client ./cmd/grpc_client

run-grpc: build
	./bin/grpc_server

run-rest: build
	./bin/rest_server


create-contact:
	./bin/grpc_client create -name "Oleg Kirka" -phone "+12345678901" -email "oleg@example.com"
# Some failing examples
create-contact-invalid-name:
	./bin/grpc_client create -name "J" -phone "+12345678901" -email "john@example.com"


create-contact-invalid-phone:
	./bin/grpc_client create -name "John Doe" -phone "123456" -email "john@example.com"

create-contact-invalid-email:
	./bin/grpc_client create -name "John Doe" -phone "+12345678901" -email "invalid-email"
#----------

list-contacts:
	./bin/grpc_client list

get-contact:
	./bin/grpc_client get -id $(ID)

update-contact:
	./bin/grpc_client update -id $(ID) -name "Vitya Shpala" -phone "+12345678902" -email "vitua@example.com"

delete-contact:
	./bin/grpc_client delete -id $(ID)

# PostgreSQL setup
setup-db:
	PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE contacts;"
	PGPASSWORD=postgres psql -h localhost -U postgres -d contacts -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

# HTTP client
http-create:
	curl -X POST http://localhost:8080/api/v1/contacts -d '{"name":"John Doe","phone":"+12345678901","email":"john@example.com"}' -H "Content-Type: application/json"

http-create-invalid:
	curl -X POST http://localhost:8080/api/v1/contacts -d '{"name":"J","phone":"123456","email":"invalid-email"}' -H "Content-Type: application/json"

http-list:
	curl -X GET http://localhost:8080/api/v1/contacts

http-get:
	curl -X GET http://localhost:8080/api/v1/contacts/$(ID)

http-update:
	curl -X PUT http://localhost:8080/api/v1/contacts/$(ID) -d '{"name":"Jane Doe","phone":"+12345678902","email":"jane@example.com"}' -H "Content-Type: application/json"

http-delete:
	curl -X DELETE http://localhost:8080/api/v1/contacts/$(ID)

clean:
	rm -rf bin/
	rm -rf gen/
