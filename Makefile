.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: build
build:
	go build -o server ./cmd/server
	go build -o shorten ./cmd/cli

.PHONY: server
server:
	go run ./cmd/server

.PHONY: cli
cli:
	go run ./cmd/cli

.PHONY: clean
clean:
	rm -f server shortener coverage.out coverage.html
	go mod tidy
	go mod vendor

.PHONY: rebuild
rebuild: clean build

.PHONY: deploy
deploy:
	gcloud app deploy app.yaml dispatch.yaml

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -cover ./...

.PHONY: html
html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor
