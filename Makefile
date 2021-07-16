.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: build
build:
	go build -o server ./cmd/server
	go build -o shortener ./cmd/cli

.PHONY: server
server:
	go run ./cmd/server

.PHONY: clean
clean:
	rm -f server shortener coverage.out coverage.html
	go mod tidy
	go mod vendor

.PHONY: rebuild
rebuild: clean build

.PHONY: deploy
deploy:
	gcloud app deploy app.yaml

.PHONY: shortener
shortener: rebuild
	./shortener

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
