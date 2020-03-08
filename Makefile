.DEFAULT_GOAL := default
DIR_NAME := `basename ${PWD}`
COVERAGE_FILE := cover.out
COVERAGE_HTML := cover.html

test:
	go test -cover ./...

coverage:
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)

dev-deps:
	go get -u golang.org/x/lint/golint

deps:
	go mod download

build:
	go build -o server ./cmd/server/main.go

install:
	go install

check:
	golangci-lint run --enable-all -D dupl ./...

clean:
	go clean ./cmd/server/main.go
	@rm -f ./server
	rm -f $(NAME) $(COVERAGE_FILE) $(COVERAGE_HTML)

default:
	make check
	make build

link:
	ln -sv `pwd` $(ALFRED_WORKFLOW_DIR)/

unlink:
	rm -i $(ALFRED_WORKFLOW_DIR)/$(DIR_NAME)

.PHONY: build install check clean test
