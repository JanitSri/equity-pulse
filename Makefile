BINARY_NAME=main
TEST_COVERAGE_REPORT=coverage.out

build:
	@go build -o ./bin/${BINARY_NAME} main.go

run: build
	@./bin/${BINARY_NAME}

install: 
	@go install	

tidy:
	@go mod tidy

clean: tidy
	@go clean

mockery: 
	docker run --rm -v "$$PWD":/src -w /src vektra/mockery ${ARGS}  # make mockery ARGS="--name StockDataProvider --dir /src/net"

test: 
	@go test ${ARGS}  # make test ARGS='-v ./... -tags="!ignoretest"'
	
test-report: 
	@go test -v -coverprofile=${TEST_COVERAGE_REPORT} ${ARGS}  # make test-report ARGS="./..."
	@go tool cover -html=${TEST_COVERAGE_REPORT}

test-fuzz: 
	@go test -fuzz=. -fuzztime=10s ${PACKAGE}  # make test-fuzz PACKAGE="./service"

compose-up:
	@docker-compose -f docker-compose.yaml up -d

compose-down:
	@docker-compose -f docker-compose.yaml down

stock: install
	equity-pulse stock $(ARGS)  # make stock ARGS="--help"
