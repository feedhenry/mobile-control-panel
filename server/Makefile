# Copyright 2016 The Kubernetes Authors.
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

all: images

# Some env vars that devs might find useful:
#  GOFLAGS      : extra "go build" flags to use - e.g. -v   (for verbose)
#  TEST_DIRS=   : only run the unit tests from the specified dirs
#  UNIT_TESTS=  : only run the unit tests matching the specified regexp

# Define some constants
#######################
ROOT           = $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
BINDIR        ?= bin
SCRIPT_DIR     ?= hack
COVERAGE      ?= $(CURDIR)/coverage.html
PKG     = github.com/feedhenry/mobile-control-panel/server
TOP_SRC_DIRS   = cmd pkg
SRC_DIRS       = $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*.go \
                   -exec dirname {} \\; | sort | uniq")
TEST_DIRS     ?= $(shell sh -c "find $(TOP_SRC_DIRS) -name \\*_test.go \
                   -exec dirname {} \\; | sort | uniq")
VERSION       ?= $(shell git describe --tags --always --abbrev=7 --dirty)
ifeq ($(shell uname -s),Darwin)
STAT           = stat -f '%c %N'
else
STAT           = stat -c '%Y %n'
endif
NEWEST_GO_FILE = $(shell find $(SRC_DIRS) -name \*.go -exec $(STAT) {} \; \
                   | sort -r | head -n 1 | sed "s/.* //")
TYPES_FILES    = $(shell find pkg/apis -name types.go)
GO_VERSION     = 1.7.3

PLATFORM?=linux
ARCH?=amd64

GO_BUILD       = env GOOS=$(PLATFORM) GOARCH=$(ARCH) go build -i $(GOFLAGS) \
                   -ldflags "-X $(PKG)/pkg.VERSION=$(VERSION)"
BASE_PATH      = $(ROOT:/src/github.com/feedhenry/mobile-control-panel/server/=)


ifdef UNIT_TESTS
	UNIT_TEST_FLAGS=-run $(UNIT_TESTS) -v
endif

NON_VENDOR_DIRS = $(shell glide nv)

# This section builds the output binaries.
# Some will have dedicated targets to make it easier to type, for example
# "apiserver" instead of "bin/apiserver".
#########################################################################
build: .generate_files server

server: $(BINDIR)/server
$(BINDIR)/server: .generate_files $(NEWEST_GO_FILE)
$(BINDIR)/server: $(NEWEST_GO_FILE)
	$(GO_BUILD) -o $@ $(PKG)/cmd/server

# This section contains the code generation stuff
#################################################
.generate_exes: $(BINDIR)/defaulter-gen \
                $(BINDIR)/deepcopy-gen \
                $(BINDIR)/conversion-gen \
                $(BINDIR)/client-gen \
                $(BINDIR)/lister-gen \
                $(BINDIR)/informer-gen
	touch $@

$(BINDIR)/defaulter-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/defaulter-gen

$(BINDIR)/deepcopy-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/deepcopy-gen

$(BINDIR)/conversion-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/conversion-gen

$(BINDIR)/client-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/client-gen

$(BINDIR)/lister-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/lister-gen

$(BINDIR)/informer-gen:
	go build -o $@ k8s.io/kubernetes/cmd/libs/go2idl/informer-gen

# Regenerate all files if the gen exes changed or any "types.go" files changed
.generate_files: .generate_exes $(TYPES_FILES)
	# Generate defaults
	$(BINDIR)/defaulter-gen \
		--v 1 --logtostderr \
		--go-header-file "hack/boilerplate.go.txt" \
		--input-dirs "$(PKG)/pkg/apis/mobile" \
		--input-dirs "$(PKG)/pkg/apis/mobile/v1alpha1" \
	  	--extra-peer-dirs "$(PKG)/pkg/apis/mobile" \
		--extra-peer-dirs "$(PKG)/pkg/apis/mobile/v1alpha1" \
		--output-file-base "zz_generated.defaults"
	# Generate deep copies
	$(BINDIR)/deepcopy-gen \
		--v 1 --logtostderr \
		--go-header-file "hack/boilerplate.go.txt" \
		--input-dirs "$(PKG)/pkg/apis/mobile" \
		--input-dirs "$(PKG)/pkg/apis/mobile/v1alpha1" \
		--bounding-dirs "github.com/feedhenry/mobile-control-panel/server" \
		--output-file-base zz_generated.deepcopy
	# Generate conversions
	$(BINDIR)/conversion-gen \
		--v 1 --logtostderr \
		--go-header-file "hack/boilerplate.go.txt" \
		--input-dirs "$(PKG)/pkg/apis/mobile" \
		--input-dirs "$(PKG)/pkg/apis/mobile/v1alpha1" \
		--output-file-base zz_generated.conversion
	# generate all pkg/client contents
	$(SCRIPT_DIR)/update-codegen.sh

# Util targets
##############
.PHONY: verify verify-client-gen 
verify: .generate_files verify-client-gen
	@echo Running gofmt:
	@gofmt -l -s $(TOP_SRC_DIRS) > .out 2>&1 || true
	@bash -c '[ "`cat .out`" == "" ] || \
	  (echo -e "\n*** Please 'gofmt' the following:" ; cat .out ; echo ; false)'
	@rm .out

verify-client-gen: .generate_files
	$(SCRIPT_DIR)/verify-codegen.sh

format:
	gofmt -w -s $(TOP_SRC_DIRS)

test: build test-unit

test-unit: build
	@echo Running tests:
	go test -cover $(UNIT_TEST_FLAGS) \
	  $(addprefix $(PKG)/,$(TEST_DIRS))

test-integration: build
	@echo Starting etcd:
	docker run --rm -d -p 2379:2379 --name mobile-etcd quay.io/coreos/etcd:v3.1.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379
	@echo Running integration tests:
	go test \
	  $(addprefix $(PKG)/,test/integration)
	docker stop mobile-etcd

clean: clean-bin

clean-bin:
	rm -rf $(BINDIR)
	rm -f .generate_exes

# Building Docker Images for our executables
############################################
images: apiserver-image

apiserver-image: artifacts/image/Dockerfile $(BINDIR)/server build
	mkdir -p tmp
	cp $(BINDIR)/server tmp
	cp artifacts/image/Dockerfile tmp
	docker build -t feedhenry/mobile-apiserver:$(VERSION) tmp
	docker tag feedhenry/mobile-apiserver:$(VERSION) feedhenry/mobile-apiserver:latest
	rm -rf tmp
