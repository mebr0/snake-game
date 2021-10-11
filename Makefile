.PHONY:
.DEFAULT_GOAL := build

lint:
	golangci-lint run

format:
	goimports -w .

test:
	go test --short -coverprofile=cover.out -v ./...

cover: test
	go tool cover -func=cover.out

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
