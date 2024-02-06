.PHONY:
	build move-to-bin install test lint dep test-ci latest-release

dep:
	go mod vendor

build:
	@echo "Building the binary..."
	@go build -v -o diff-check .
	@echo "Done."

move-to-bin:
	@echo "Moving binary to /usr/local/bin...requires sudo"
	@sudo cp diff-check /usr/local/bin
	@echo "Done."

install:
	make build && make move-to-bin

lint: install-static-check lint vet

vet:
	go vet ./...

install-static-check:
	go install honnef.co/go/tools/cmd/staticcheck@latest

lint-static-check:
	staticcheck ./...

lint: vet lint-static-check

test:
	go test ./...

test-ci:
	gotestsum --format testname --junitfile pkg/results.xml -- -coverprofile=cover.out ./...

latest-release:
	curl -sL https://api.github.com/repos/automoto/diff-check/releases/latest | jq -r '.assets[].browser_download_url'