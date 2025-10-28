SHELL = /bin/bash

SCRIPT_DIR      = $(shell pwd)/scripts

MAJOR_VERSION   = $(shell cat VERSION)
GIT_VERSION     = $(shell git log -1 --pretty=format:%h)
GIT_NOTES       = $(shell git log -1 --oneline)

unit-test-local:
	GO111MODULE=on go test -gcflags=-l -v `go list .` -run="^(TestUnmarshal|TestForAIGW).*"

coverage-local:
	bash ${SCRIPT_DIR}/report.sh

.PHONY: unit-test coverage-local
