.PHONY: lint
lint:
	golangci-lint run .

.PHONY: test
test:
	go test --cover ./...

.PHONY: test-flakyness
test-flakyness:
	parallel -N0 go test ./... ::: {1..50}
