BINARY_NAME=main

build:
	@go build -o ./bin/${BINARY_NAME} main.go

run: build
	@./bin/${BINARY_NAME}

clean:
	@go clean
	@rm -rf ./bin

install: 
	@go install	

company: install
	equity-pulse company $(ARGS)