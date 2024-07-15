build:
	@go build -o bin/green-api cmd/green-api/main.go

test:
	@go test -v ./...

run: build
	@./bin/green-api