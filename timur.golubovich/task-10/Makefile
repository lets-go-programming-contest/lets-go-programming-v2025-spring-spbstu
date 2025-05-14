preinstall:
	git submodule add https://github.com/bufbuild/protovalidate.git api/third_party/protovalidate
	git submodule add https://github.com/googleapis/googleapis.git api/third_party/googleapis
	go install github.com/bufbuild/buf/cmd/buf@v1.30.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2

build-server:
	go build -o build/phonebook ./cmd/phonebook

build-client:
	go build -o build/phonebook_grpc_client ./cmd/phonebook_grpc_client

build-http-server:
	go build -o build/phonebook_rest_api ./cmd/phonebook_rest_api

build: build-server build-client build-http-server

generate:
	go generate ./api/proto

run-server: build-server
	./build/phonebook

run-rest-api:
	./bin/phonebook_rest_api

setup-db:
	PGPASSWORD=postgres psql -h localhost -U postgres -c "CREATE DATABASE phonebook;"

clean:
	rm -rf build/
	rm -rf gen/
	rm -rf docs/

GRPC_CLIENT_PATH = ./build/phonebook_grpc_client
CURL = curl -s
HTTP_BASE_URL = http://localhost:8080/v1
DEFAULT_ID = 237aab70-905a-4e38-ad3a-3d608f0dd5e1
ID ?= $(DEFAULT_ID)
NAME?=Dill Doe
PHONE?=+79142468804


grpc_create_contact:
	$(GRPC_CLIENT_PATH) create-contact --name $(NAME) --phone $(PHONE) | jq

grpc_create_contact_invalid:
	$(GRPC_CLIENT_PATH) create-contact --name "J" --phone "+1234567890" | jq

grpc_contacts_list:
	$(GRPC_CLIENT_PATH) list-contacts | jq

grpc_get_contact:
	$(GRPC_CLIENT_PATH) get-contact --id "$(ID)" | jq

grpc_update_contact:
	$(GRPC_CLIENT_PATH) update-contact --id "$(ID)" --name "$(NAME)" --phone "$(PHONE)" | jq

grpc_delete_contact:
	$(GRPC_CLIENT_PATH) delete-contact --id "$(ID)" | jq


http_create_contact:
	$(CURL) -X POST -H "Content-Type: application/json" \
		-d '{"name":"$(NAME)","phone":"$(PHONE)"}' \
		$(HTTP_BASE_URL)/contacts | jq

http_contacts_list:
	$(CURL) $(HTTP_BASE_URL)/contacts | jq

http_get_contact:
	$(CURL) $(HTTP_BASE_URL)/contacts/$(ID) | jq

http_update_contact:
	$(CURL) -X PUT -H "Content-Type: application/json" \
		-d '{"name":"$(NAME)","phone":"$(PHONE)"}' \
		$(HTTP_BASE_URL)/contacts/$(ID) | jq

http_delete_contact:
	$(CURL) -X DELETE $(HTTP_BASE_URL)/contacts/$(ID) | jq
