CMD_PACKAGE := github.com/sters/jij/cmd/jij
GO_ENV := GO111MODULE=on CGO_ENABLED=0
BUILD_DIR := ./build

.PHONY: init tidy test build install
init: 
	${GO_ENV} go mod init
tidy: 
	${GO_ENV} go mod tidy
test: 
	${GO_ENV} go test -v ./...
build: 
	${GO_ENV} go build -o $(BUILD_DIR)/jij cmd/main.go
install:
	${GO_ENV} go install ${CMD_PACKAGE}
run:
	${GO_ENV} go run ${CMD_PACKAGE}