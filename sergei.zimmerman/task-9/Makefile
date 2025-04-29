OUT_DIR := build/bin
GO := go
CMD_DIR := cmd
BASH := bash

.PHONY: all
all: $(OUT_DIR)/phonebook-server

$(OUT_DIR):
	mkdir -p $(OUT_DIR)

$(OUT_DIR)/phonebook-server: $(OUT_DIR)
	$(GO) build -o $@ $(CMD_DIR)/phonebook-server/main.go

.PHONY: $(OUT_DIR)/phonebook-server

.PHONY: test

test: $(OUT_DIR)/phonebook-server
	$(BASH) scripts/test_phonebook.sh --config config/config.yaml --binary $(OUT_DIR)/phonebook-server
