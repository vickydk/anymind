#!/bin/bash

export REPO_NAME=anymind

build:
	@echo "${NOW} Building Service"
	@go build -o ./bin/${REPO_NAME}-grpc cmd/app/main.go

docker-build:
	@ echo "Building anymind image"
	@ docker build -f Dockerfile -t anymind .

run-app: build-app
	@./bin/${REPO_NAME}-app

compose: compose_down
	go mod vendor
	docker-compose up --build --remove-orphans

compose_down:
	docker-compose down

run_benchmark:
	@cat benchmark/vegeta_requests.list | vegeta attack -duration=5s -rate=10 | vegeta report --type=text

staticcheck_install:
	go install honnef.co/go/tools/cmd/staticcheck@latest

staticcheck:
	staticcheck -go 1.18 ./...

linter_install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1

linter:
	golangci-lint run -v --disable=typecheck --disable=structcheck --timeout=5m --go=1.18