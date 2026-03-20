build:
	@go build -o ./bin/fs ./cmd/server/*.go

run: build
	@./bin/fs
	
test:
	@go test ./... -v
