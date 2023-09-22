start:
	go run ./cmd/mygram-api/main.go

build:
	go build -v -o bin/main ./cmd/mygram-api/main.go

.PHONY: test
test:
	go test -v ./test