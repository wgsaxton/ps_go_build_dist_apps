# Dist Apps Makefile
SHELL=/bin/bash
GIT_TAG = $(shell git describe --tags --abbrev=0)

all:

build_logservice_container:
	docker build --tag logservice:${GIT_TAG} --file app/cmd/logservice/Dockerfile .
	docker image tag logservice:${GIT_TAG} logservice:latest