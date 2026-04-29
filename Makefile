SHELL = /bin/bash
VERSION ?= "latest"
build_dir := ./.build

.PHONY: clean
clean:
	rm -rf ${build_dir}

.PHONY: clean-cache
clean-cache:
	go clean -cache -modcache -i -r

.PHONY: build
build: clean generate discover
	@echo "Go version: $(GO_VERSION)"
	@go mod download
	@go build $(GO_BUILD_FLAGS) -o ${build_dir}/server/server ./cmd/server/

.PHONY: mod
mod:
	go mod download && go mod tidy && go mod verify

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint: mod vet fmt
	echo "Running golangci-lint..."
	@golangci-lint run -c .golangci.yml --timeout=5m

.PHONY: test
test:
	go test -race -count=1 ./...

.PHONY: service-from-static
service-from-static:
	@echo "Generating new static service..."
	@go run github.com/mockzilla/mockzilla/v2/cmd/gen/service -type static -name=$(name) -output=pkg/$(name)

.PHONY: service
service:
	@echo "Generating new OpenAPI service..."
	@go run github.com/mockzilla/mockzilla/v2/cmd/gen/service -type openapi -name=$(name) -output=pkg/$(name)

.PHONY: discover
discover:
	@echo "Discovering services to generate service imports..."
	@go run github.com/mockzilla/mockzilla/v2/cmd/gen/discover pkg

#.PHONY: generate
generate:
	@go generate ./...
