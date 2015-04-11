.PHONY: all

GOPATH_MOUNT := -v ~/go/pkg:/go/pkg
IMAGENAME := dockompleter
TIMESTAMP := $(shell date +%Y%m%d-%H%M%S)

default: generate

all: generate

generate: build
	docker run --rm $(GOPATH_MOUNT) $(IMAGENAME)

build:
	docker build -t "$(IMAGENAME):$(TIMESTAMP)" .
	docker tag -f $(IMAGENAME):$(TIMESTAMP) $(IMAGENAME):latest

