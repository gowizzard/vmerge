# Here you can reformat, check or build the binary
fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./... --timeout 5m