APP_NAME := $(shell basename "$(CURDIR)")
SHELL := /bin/bash

default: all

so-builder:
	docker build --rm -t $(APP_NAME)/builder . -f ./Dockerfile-libplctag

so-lib: so-builder
	docker run --rm -v "$(CURDIR)/libs:/artifact" -v "$(CURDIR):/go/src/github.com/brcz/go-plctag" $(APP_NAME)/builder

test-builder:
	docker build --rm -t $(APP_NAME)/tbuilder . -f ./Dockerfile-test

test-lib: test-builder
	docker run --rm -v "$(CURDIR)/bin:/artifact" -v "$(CURDIR):/go/src/github.com/brcz/go-plctag" $(APP_NAME)/tbuilder

sync:
	rsync -zvr  ./third_party/libplctag/src/ brcz@192.168.99.100:/home/brcz/go/src/github.com/go-plctag/third_party/libplctag/src