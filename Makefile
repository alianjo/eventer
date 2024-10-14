GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin
GOFILES=$(wildcard *.go)
IMG ?= eventer:latest

build:
	go build -o bin/main $(GOFILES)

docker-build:
	docker build -t ${IMG} .

docker-push:
	docker push ${IMG}

docker-run:
	docker run ${IMG}

run: build
	bin/main

all: docker-build docker-push docker-run

.PHONY: build docker-build docker-push docker-run run all


deploy:
	kubectl create -f manifests/