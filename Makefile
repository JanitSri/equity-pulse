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

stock: install
	equity-pulse stock $(ARGS)

tidy:
	@go mod tidy
