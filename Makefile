#
# Copyright 2021 The Sigstore Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all test clean clean-gen lint gosec ko ko-local cross-cli gocovmerge

all: rekor-cli rekor-server ## Build all binaries (rekor-cli and rekor-server)

include Makefile.swagger

OPENAPIDEPS = openapi.yaml $(shell find pkg/types -iname "*.json")
SRCS = $(shell find cmd -iname "*.go") $(shell find pkg -iname "*.go"|grep -v pkg/generated) pkg/generated/restapi/configure_rekor_server.go $(SWAGGER_GEN)
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)
BIN_DIR := $(abspath $(ROOT_DIR)/bin)
FUZZ_DURATION ?= 10s

PROJECT_ID ?= projectsigstore
RUNTIME_IMAGE ?= gcr.io/distroless/static
# Set version variables for LDFLAGS
GIT_VERSION ?= $(shell git describe --tags --always --dirty)
GIT_HASH ?= $(shell git rev-parse HEAD)
DATE_FMT = +%Y-%m-%dT%H:%M:%SZ
SOURCE_DATE_EPOCH ?= $(shell git log -1 --pretty=%ct)
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif
GIT_TREESTATE = "clean"
DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(DIFF), 1)
    GIT_TREESTATE = "dirty"
endif

KO_PREFIX ?= gcr.io/projectsigstore
export KO_DOCKER_REPO=$(KO_PREFIX)
REKOR_YAML ?= rekor-$(GIT_VERSION).yaml
GHCR_PREFIX ?= ghcr.io/sigstore/rekor
GOBIN ?= $(shell go env GOPATH)/bin

# Binaries
SWAGGER := $(TOOLS_BIN_DIR)/swagger
GO-FUZZ-BUILD := $(TOOLS_BIN_DIR)/go-fuzz-build
GOCOVMERGE := $(TOOLS_BIN_DIR)/gocovmerge

REKOR_LDFLAGS=-X sigs.k8s.io/release-utils/version.gitVersion=$(GIT_VERSION) \
              -X sigs.k8s.io/release-utils/version.gitCommit=$(GIT_HASH) \
              -X sigs.k8s.io/release-utils/version.gitTreeState=$(GIT_TREESTATE) \
              -X sigs.k8s.io/release-utils/version.buildDate=$(BUILD_DATE)

CLI_LDFLAGS=$(REKOR_LDFLAGS)
SERVER_LDFLAGS=$(REKOR_LDFLAGS)

Makefile.swagger: $(SWAGGER) $(OPENAPIDEPS) ## Generate Swagger code and Makefile
	$(SWAGGER) validate openapi.yaml
	$(SWAGGER) generate client -f openapi.yaml -q -r COPYRIGHT.txt -t pkg/generated --additional-initialism=TUF --additional-initialism=DSSE
	$(SWAGGER) generate server -f openapi.yaml -q -r COPYRIGHT.txt -t pkg/generated --exclude-main -A rekor_server --flag-strategy=pflag --default-produces application/json --additional-initialism=TUF --additional-initialism=DSSE
	@echo "# This file is generated after swagger runs as part of the build; do not edit!" > Makefile.swagger
	@echo "SWAGGER_GEN=`find pkg/generated/client pkg/generated/models pkg/generated/restapi -iname '*.go' | grep -v 'configure_rekor_server' | sort -d | tr '\n' ' ' | sed 's/ $$//'`" >> Makefile.swagger;

lint: ## Run golangci-lint checks
	$(GOBIN)/golangci-lint run -v ./...

gosec: ## Run gosec security scanner
	$(GOBIN)/gosec ./...

rekor-cli: $(SRCS) ## Build the rekor command-line interface
	CGO_ENABLED=0 go build -trimpath -ldflags "$(CLI_LDFLAGS)" -o rekor-cli ./cmd/rekor-cli

rekor-server: $(SRCS) ## Build the rekor server
	CGO_ENABLED=0 go build -trimpath -ldflags "$(SERVER_LDFLAGS)" -o rekor-server ./cmd/rekor-server

backfill-index: $(SRCS) ## Build the backfill-index utility
	CGO_ENABLED=0 go build -trimpath -ldflags "$(SERVER_LDFLAGS)" -o backfill-index ./cmd/backfill-index

test: ## Run all tests
	go test ./...

gocovmerge: $(GOCOVMERGE) ## Merge coverage profiles

clean: ## Remove built binaries and artifacts
	rm -rf dist
	rm -rf hack/tools/bin
	rm -rf rekor-cli rekor-server
	rm -f *fuzz.zip

clean-gen: clean ## Clean generated files and swagger code
	rm -rf $(SWAGGER_GEN)

up: ## Start services using docker-compose
	docker-compose -f docker-compose.yml build --build-arg SERVER_LDFLAGS="$(SERVER_LDFLAGS)"
	docker-compose -f docker-compose.yml up

debug: ## Start services in debug mode
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml build --build-arg SERVER_LDFLAGS="$(SERVER_LDFLAGS)" rekor-server-debug
	docker-compose -f docker-compose.yml -f docker-compose.debug.yml up rekor-server-debug

ko: ## Build and publish container images using ko
	# rekor-server
	LDFLAGS="$(SERVER_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	KO_DOCKER_REPO=$(KO_PREFIX)/rekor-server ko resolve --bare \
		--platform=all --tags $(GIT_VERSION) --tags $(GIT_HASH) \
		--image-refs rekorServerImagerefs --filename config/ > $(REKOR_YAML)

	# rekor-cli
	LDFLAGS="$(CLI_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	ko publish --base-import-paths \
		--platform=all --tags $(GIT_VERSION) --tags $(GIT_HASH) \
		--image-refs rekorCliImagerefs github.com/Morrison76/rekor/cmd/rekor-cli

	# backfill-index
	LDFLAGS="$(SERVER_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	ko publish --base-import-paths \
		--platform=all --tags $(GIT_VERSION) --tags $(GIT_HASH) \
		--image-refs bIndexImagerefs github.com/Morrison76/rekor/cmd/backfill-index

deploy: ## Deploy to Kubernetes using ko
	LDFLAGS="$(SERVER_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) ko apply -f config/

e2e: ## Build and prepare end-to-end tests
	go test -c -tags=e2e ./tests
	go test -c -tags=e2e ./pkg/pki/x509
	go test -c -tags=e2e ./pkg/pki/tuf
	go test -c -tags=e2e ./pkg/types/rekord

sign-keyless-ci: ko ## Sign container images using keyless signing
	cosign sign --yes -a GIT_HASH=$(GIT_HASH) $(KO_DOCKER_REPO)/rekor-server:$(GIT_HASH)
	cosign sign --yes -a GIT_HASH=$(GIT_HASH) $(KO_DOCKER_REPO)/rekor-cli:$(GIT_HASH)

ko-local: ## Build container images locally using ko
	KO_DOCKER_REPO=ko.local LDFLAGS="$(SERVER_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	ko publish --base-import-paths \
		--tags $(GIT_VERSION) --tags $(GIT_HASH) --image-refs rekorImagerefs \
		github.com/Morrison76/rekor/cmd/rekor-server

	KO_DOCKER_REPO=ko.local LDFLAGS="$(CLI_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	ko publish --base-import-paths \
		--tags $(GIT_VERSION) --tags $(GIT_HASH) --image-refs cliImagerefs \
		github.com/Morrison76/rekor/cmd/rekor-cli

	KO_DOCKER_REPO=ko.local LDFLAGS="$(SERVER_LDFLAGS)" GIT_HASH=$(GIT_HASH) GIT_VERSION=$(GIT_VERSION) \
	ko publish --base-import-paths \
		--tags $(GIT_VERSION) --tags $(GIT_HASH) --image-refs indexImagerefs \
		github.com/Morrison76/rekor/cmd/backfill-index

fuzz: ## Run fuzz tests for a configured duration
	go test -fuzz FuzzCreateEntryIDFromParts -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzGetUUIDFromIDString -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzGetTreeIDFromIDString -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzPadToTreeIDLen -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzReturnEntryIDString -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzTreeID -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzValidateUUID -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzValidateTreeID -fuzztime $(FUZZ_DURATION) ./pkg/sharding
	go test -fuzz FuzzValidateEntryID -fuzztime $(FUZZ_DURATION) ./pkg/sharding

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(GO-FUZZ-BUILD): $(TOOLS_DIR)/go.mod ## Build go-fuzz-build tool
	cd $(TOOLS_DIR); go build -trimpath -tags=tools -o $(TOOLS_BIN_DIR)/go-fuzz-build github.com/dvyukov/go-fuzz/go-fuzz-build

$(SWAGGER): $(TOOLS_DIR)/go.mod ## Build swagger tool
	cd $(TOOLS_DIR); go build -trimpath -tags=tools -o $(TOOLS_BIN_DIR)/swagger github.com/go-swagger/go-swagger/cmd/swagger

$(GOCOVMERGE): $(TOOLS_DIR)/go.mod ## Build gocovmerge tool
	cd $(TOOLS_DIR); go build -trimpath -tags=tools -o $(TOOLS_BIN_DIR)/gocovmerge github.com/wadey/gocovmerge

##################
# help
##################

help: ## Display this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

include release/release.mk