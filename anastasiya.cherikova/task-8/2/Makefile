.PHONY: all prod debug clean

prod:
	@echo "Building production version..."
	@go build -o main

debug:
	@echo "Building debug version..."
	@go build -tags debug -o main_debug

clean:
	@rm -f main main_debug

run-prod: prod
	@./main

run-debug: debug
	@./main_debug