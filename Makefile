.PHONY: clean
clean:
	@rm -rf advent-of-code-go ./input

.PHONY: test
test:
	@go test ./...
