run: build
	@./bin/skeleton

build:
	go build -o bin/skeleton ./cmd/skeleton/
	@echo "Build completed!"