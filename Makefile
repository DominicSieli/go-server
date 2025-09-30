build:
	@go build -o bin/go-server-template cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-server-template