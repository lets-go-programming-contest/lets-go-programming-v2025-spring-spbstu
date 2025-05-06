# Get all contacts
curl -v -s "http://localhost:8080/contacts" | jq .

# Get contact by ID
curl -v -s "http://localhost:8080/contacts/$CONTACT_ID" | jq .

# Create contact
curl -v -s -X POST "http://localhost:8080/contacts" \
    -H "Content-Type: application/json" \
    -d '{"name": "Bob", "phone": "+987654321"}'

# Update contact
curl -v -s -X PUT "http://localhost:8080/contacts/$CONTACT_ID" \
    -H "Content-Type: application/json" \
    -d '{"name": "Alice", "phone": "+123456789"}'

# Delete contact
curl -v -s -X DELETE "http://localhost:8080/contacts/$CONTACT_ID"
