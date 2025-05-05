build-server:
	go build -o bin/contact_manager ./cmd/contact_manager

build-client:
	go build -o bin/contact_manager_grpc_client ./cmd/contact_manager_grpc_client

build-http-server:
	go build -o bin/contact_manager_rest_api ./cmd/contact_manager_rest_api

build: build-server build-client build-http-server

generate:
	go generate ./api/proto

run-server: build-server
	./bin/contact_manager

setup-db:
	PGPASSWORD=$(PGPASSWORD) psql -h localhost -U postgres -c "CREATE DATABASE contact_manager;"

clean:
	rm -rf bin/
	rm -rf gen/


GRPC_CLIENT_PATH = ./bin/contact_manager_grpc_client
CURL = curl -s
HTTP_BASE_URL = http://localhost:8080/v1
DEFAULT_ID = 1
ID ?= $(DEFAULT_ID)
NAME?=John Doe
PHONE?=+1234567890

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
