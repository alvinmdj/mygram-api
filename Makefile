start:
	go run ./cmd/mygram-api/main.go

build:
	go build -v -o bin/mygram-api ./cmd/mygram-api/main.go

.PHONY: test
test:
	go test -v ./test