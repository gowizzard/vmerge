# Here you can reformat, check or build the binary
BINARY_NAME=vmerge

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./... --timeout 5m

build:
	go build -o ${BINARY_NAME}