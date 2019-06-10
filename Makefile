APP_NAME := $(shell basename "$(CURDIR)")
SHELL := /bin/bash

default: all

so-builder:
	docker build --rm -t $(APP_NAME)/builder . -f ./Dockerfile-libplctag

so-lib: so-builder
	docker run --rm -v "$(CURDIR)/libs:/artifact" -v "$(CURDIR):/go/src/github.com/brcz/go-plctag" $(APP_NAME)/builder

sync:
	