.PHONY: clean
clean:
	@rm -rf advent-of-code-go ./input

.PHONY: test
test:
	@go test ./...

./advent-of-code-go: ./main.go ./**/*.go

build: ./advent-of-code-go
	@go build

run: build
	@./advent-of-code-go
