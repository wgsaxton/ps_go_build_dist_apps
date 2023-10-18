# Dist Apps Makefile
SHELL=/bin/bash
GIT_TAG = $(shell git describe --tags --abbrev=0)
export REPO_ROOT := $(shell git rev-parse --show-toplevel)


all:

build_logservice_container:
	docker build --tag logservice:${GIT_TAG} --file app/cmd/logservice/Dockerfile .
	docker image tag logservice:${GIT_TAG} logservice:latest

build_registryservice_container:
	docker build --tag registryservice:${GIT_TAG} --file app/cmd/registryservice/Dockerfile .
	docker image tag registryservice:${GIT_TAG} registryservice:latest

build_binaries:
	app/build.sh