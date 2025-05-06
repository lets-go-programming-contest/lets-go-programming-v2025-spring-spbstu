build:
	go build -o bin/contact_manager ./cmd/contact_manager

run: build
	./bin/contact_manager

setup-db:
	PGPASSWORD=$(PGPASSWORD) psql -h localhost -U postgres -c "CREATE DATABASE contact_manager;"

clean:
	rm -rf bin/

CURL = curl -s
HTTP_BASE_URL = http://localhost:8080/v1
NAME?=John Doe
PHONE?=+1234567890

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
