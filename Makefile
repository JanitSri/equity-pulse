BINARY_NAME=main

build:
	@go build -o ./bin/${BINARY_NAME} main.go

run: build
	@./bin/${BINARY_NAME}

clean: tidy
	@go clean
	@rm -rf ./bin

install: 
	@go install	

tidy:
	@go mod tidy

compose-up:
	@docker-compose -f docker-compose.yaml up -d

compose-down:
	@docker-compose -f docker-compose.yaml down

stock: install
	equity-pulse stock $(ARGS)
