# Makefile
SHELL:=/bin/bash

APPNAME  =`basename ${PWD}`
DATE     =`date -u +%Y%m%d%H%M%S`
REVISION =`git rev-parse HEAD`
VERSION  =`git describe --abbrev=0 --tags | sed 's/v//'`

WORKDIR  = tmp
BUILDDIR = build
SUBDIRS  =`go list -e ./... | grep -v /vendor/ | grep -v binary | grep -v $(shell basename ${PWD})$$`

.PHONY: all default test build install vendor dep run clean test-local-pre test-local test-local-post

default: all
all: $(APPNAME)

$(APPNAME): build

dep:
	glide up -v -s --update-vendored

test:
	./test.sh

# test main with coverage profile
test-main:
	mkdir -p $(BUILDDIR)
	go vet -a . || exit 1
	go test -v \
		-covermode atomic \
		-coverprofile build/coverage.out \
		-trace build/trace.out \
		-race \
		-ginkgo.v -ginkgo.randomizeAllSpecs -ginkgo.progress

# test all subdir packages excluding vendor
test-packages:
	mkdir -p $(BUILDDIR)
	for var in $(SUBDIRS) ; do \
		echo "testing package $$var" ; \
		go vet -a $$var || exit 1 ; \
		go test -v \
			-covermode atomic \
			-race \
			$$var || exit 1 ; \
	done

# test with local go-common changes
test-local: test-local-pre
	@$(MAKE) --no-print-directory test-packages || exit 0
	@$(MAKE) --no-print-directory test-main || exit 0
	@$(MAKE) --no-print-directory test-local-post

# move vendor directories in and out of place
# don't error out here if we cannot move it due to previous test failure
test-local-pre:
	@mv vendor _vendor 2> /dev/null || exit 0
	@mv $(GOPATH)/src/github.com/bitgaming/go-common/vendor \
		$(GOPATH)/src/github.com/bitgaming/go-common/_vendor 2> /dev/null || exit 0

test-local-post:
	@mv _vendor vendor 2> /dev/null || exit 0
	@mv $(GOPATH)/src/github.com/bitgaming/go-common/_vendor \
		$(GOPATH)/src/github.com/bitgaming/go-common/vendor 2> /dev/null || exit 0

# linker sets variables in version.go here
# Built, Revision, AppName, Version
build: *.go
	@echo "using $(shell go version)"
	CGO_ENABLED=0 go build -ldflags \
		 "-w -s \
		 -X main.Built=$(DATE) \
		 -X main.Revision=$(REVISION) \
		 -X main.AppName=$(APPNAME) \
		 -X main.Version=$(VERSION)" \
		 -o $(APPNAME) \
		 -tags netgo -a
	# test print version information added by linker
	./$(APPNAME) -version

install: $(APPNAME)
	go install

run: $(APPNAME)
	./$(APPNAME)

clean:
	\rm -rf $(APPNAME) $(WORKDIR) $(BUILDDIR)
