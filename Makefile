build:
	@go build -o bin/check-euro
run: build
	@./bin/check-euro
test:
	@go test -v ./...