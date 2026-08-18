SHELL := /bin/sh

.PHONY: help fmt lint test test-cover test-race test-benchmark vet staticcheck bench 

help:
	@printf '%s\n' \
		'make fmt            format Go sources' \
		'make lint           run linter' \
		'make test           run unit tests' \
		'make test-cover     run tests and show coverage' \
		'make test-race      run tests with race detector' \
		'make test-benchmark run benchmark' \
		'make vet            run go vet' \
		'make staticcheck    run staticcheck when installed' \
		'make bench          run benchmarks' \

fmt:
	gofmt -s -w $$(find . -name '*.go' -type f)

lint:
	golangci-lint run ./...

test:
	go test ./...

test-cover:
	go test --cover ./...

test-race:
	go test -race ./...

test-benchmark:
	go test ./internal/cloc -run . -bench . -benchmem

vet:
	go vet ./...

staticcheck:
	@command -v staticcheck >/dev/null 2>&1 || { echo 'staticcheck is not installed' >&2; exit 1; }
	staticcheck ./...

bench:
	go test -bench=. -benchmem ./...
