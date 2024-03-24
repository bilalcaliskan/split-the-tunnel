LOCAL_BIN := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/.bin

DEFAULT_GO_TEST_CMD ?= go test ./... -race -p 1 -covermode=atomic
DEFAULT_GO_RUN_ARGS ?= ""

GOLANGCI_LINT_VERSION := latest
REVIVE_VERSION := v1.3.4
MOCKERY_VERSION := v2.39.1
# Variables
PROTO_DIR := ./api/proto
GO_PROTO_DIR := ./pkg/pb
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)
GO_OUT_OPTS := "paths=source_relative"
PROTOC_VERSION := latest


.PHONY: all
all: clean tools lint test build

.PHONY: clean
clean:
	rm -rf $(LOCAL_BIN)

.PHONY: pre-commit-setup
pre-commit-setup:
	#python3 -m venv venv
	#source venv/bin/activate
	#pip3 install pre-commit
	pre-commit install -c build/ci/.pre-commit-config.yaml

.PHONY: tools
tools:  mockery-install golangci-lint-install revive-install protoc_install vendor

.PHONY: mockery-install
mockery-install:
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@$(MOCKERY_VERSION)

.PHONY: golangci-lint-install
golangci-lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

.PHONY: revive-install
revive-install:
	GOBIN=$(LOCAL_BIN) go install github.com/mgechev/revive@$(REVIVE_VERSION)

# Ensure the necessary Go packages are installed
.PHONY: protoc_install
protoc_install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_VERSION)
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_VERSION)

.PHONY: lint
lint: tools run-lint

.PHONY: run-lint
run-lint: lint-golangci-lint lint-revive

.PHONY: lint-golangci-lint
lint-golangci-lint:
	$(info running golangci-lint...)
	$(LOCAL_BIN)/golangci-lint -v run ./... || (echo golangci-lint returned an error, exiting!; sh -c 'exit 1';)

.PHONY: lint-revive
lint-revive:
	$(info running revive...)
	$(LOCAL_BIN)/revive -formatter=stylish -config=build/ci/.revive.toml -exclude ./vendor/... ./... || (echo revive returned an error, exiting!; sh -c 'exit 1';)

.PHONY: upgrade-deps
upgrade-deps: vendor
	for item in `grep -v 'indirect' go.mod | grep '/' | cut -d ' ' -f 1`; do \
		echo "trying to upgrade direct dependency $$item" ; \
		go get -u $$item ; \
  	done
	go mod tidy
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor: tidy
	go mod vendor

.PHONY: test
test: vendor
	$(info starting the test for whole module...)
	$(DEFAULT_GO_TEST_CMD) -coverprofile=coverage.txt || (echo an error while testing, exiting!; sh -c 'exit 1';)

#.PHONY: test-unit
#test-unit: vendor
#	$(info starting the unit test for whole module...)
#	$(DEFAULT_GO_TEST_CMD) -tags "unit" -coverprofile=unit_coverage.txt || (echo an error while testing, exiting!; sh -c 'exit 1';)
#
#.PHONY: test-e2e
#test-e2e: vendor
#	$(info starting the e2e test for whole module...)
#	$(DEFAULT_GO_TEST_CMD) -tags "e2e" -coverprofile=e2e_coverage.txt || (echo an error while testing, exiting!; sh -c 'exit 1';)
#
#.PHONY: test-integration
#test-integration: vendor
#	$(info starting the integration test for whole module...)
#	$(DEFAULT_GO_TEST_CMD) -tags "integration" -coverprofile=integration_coverage.txt || (echo an error while testing, exiting!; sh -c 'exit 1';)

.PHONY: test-coverage
test-coverage: test
	go tool cover -html=coverage.txt -o cover.html
	open cover.html

.PHONY: build-daemon
build-daemon: vendor
	$(info building daemon binary...)
	go build -o .bin/split-the-tunnel cmd/daemon/daemon.go || (echo an error while building daemon binary, exiting!; sh -c 'exit 1';)

.PHONY: build-cli
build-cli: vendor
	$(info building cli binary...)
	go build -o ./.bin/stt-cli cmd/cli/cli.go || (echo an error while building cli binary, exiting!; sh -c 'exit 1';)

.PHONY: run-daemon
run-daemon: build-daemon
	$(info running daemon...)
	chmod +x ./.bin/split-the-tunnel
	sudo ./.bin/split-the-tunnel $(DEFAULT_GO_RUN_ARGS)

.PHONY: run-cli
run-cli: build-cli
	$(info running cli...)
	chmod +x ./.bin/stt-cli
	sudo ./.bin/stt-cli $(DEFAULT_GO_RUN_ARGS)

.PHONY: prepare-initial-project
GITHUB_USERNAME ?= $(shell read -p "Your Github username(ex: bilalcaliskan): " github_username; echo $$github_username)
PROJECT_NAME ?= $(shell read -p "'Kebab-cased' Project Name(ex: golang-cli-template): " project_name; echo $$project_name)
prepare-initial-project:
	grep -rl bilalcaliskan . --exclude={README.md,Makefile} --exclude-dir=.git --exclude-dir=.idea | xargs sed -i 's/bilalcaliskan/$(GITHUB_USERNAME)/g'
	grep -rl golang-cli-template . --exclude-dir=.git --exclude-dir=.idea | xargs sed -i 's/golang-cli-template/$(PROJECT_NAME)/g'
	echo "Please refer to *Additional nice-to-have steps* in README.md for additional features"
	echo "Cheers!"

.PHONY: generate-mocks
generate-mocks: mockery-install tidy vendor
	$(LOCAL_BIN)/mockery || (echo mockery returned an error, exiting!; sh -c 'exit 1';)

.PHONY: protogen
protogen: protoc_install
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT_OPTS):$(GO_PROTO_DIR) \
		--go-grpc_out=$(GO_OUT_OPTS):$(GO_PROTO_DIR) \
		$(PROTO_FILES)


# Clean up generated files
.PHONY: protoclean
protoclean:
	rm -f $(GO_PROTO_DIR)/*.go

.PHONY: release-daemon
release-daemon: build-daemon
	chmod +x bin/split-the-tunnel
	sudo cp bin/split-the-tunnel /usr/local/bin/split-the-tunnel
