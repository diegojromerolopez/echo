.PHONY: test lint

test:
	go test -v ./...

lint:
	go vet ./...
