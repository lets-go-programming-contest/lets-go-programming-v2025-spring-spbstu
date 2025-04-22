#!/bin/bash

set -euo pipefail
set -x

# Default values
CONFIG_FILE="config.yaml"
BINARY_PATH="phonebook-server"

# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
    --config)
        CONFIG_FILE="$2"
        shift
        ;;
    --binary)
        BINARY_PATH="$2"
        shift
        ;;
    *)
        echo "Unknown parameter passed: $1"
        exit 1
        ;;
    esac
    shift
done

# Check if the config file exists
if [[ ! -f "$CONFIG_FILE" ]]; then
    echo "Config file not found: $CONFIG_FILE"
    exit 1
fi

# Check if the binary exists
if [[ ! -f "$BINARY_PATH" ]]; then
    echo "Binary not found at: $BINARY_PATH"
    exit 1
fi

# Read config values
PORT=$(yq --raw-output '.server.port' "$CONFIG_FILE")
API_URL="http://localhost:$PORT"

# Start the server in the background
"$BINARY_PATH" --config "$CONFIG_FILE" &
SERVER_PID=$!

# Ensure server is stopped on exit
cleanup() {
    kill $SERVER_PID
}
trap cleanup EXIT

sleep 1

# Create contact
CREATE_RESPONSE=$(curl -v -s -X POST "$API_URL/contacts" \
    -H "Content-Type: application/json" \
    -d '{"name": "Bob", "phone": "+987654321"}')

CONTACT_ID=$(echo "$CREATE_RESPONSE" | jq -r '.id')

# List contacts
curl -v -s "$API_URL/contacts" | jq .

# Get contact by ID
curl -v -s "$API_URL/contacts/$CONTACT_ID" | jq .

# Update contact
curl -v -s -X PUT "$API_URL/contacts/$CONTACT_ID" \
    -H "Content-Type: application/json" \
    -d '{"name": "Alice", "phone": "+123456789"}'

# Get updated contact
curl -v -s "$API_URL/contacts/$CONTACT_ID" | jq .

# Delete contact
curl -v -s -X DELETE "$API_URL/contacts/$CONTACT_ID"

# Verify deletion
curl -v -s "$API_URL/contacts/$CONTACT_ID"
