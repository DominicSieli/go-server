build:
	@go build -o bin/go-server cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-server