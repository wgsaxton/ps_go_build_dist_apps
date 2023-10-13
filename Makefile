# Dist Apps Makefile
SHELL=/bin/bash
GIT_TAG = $(shell git describe)

all:

build_logservice_container:
	docker build --tag logservice:v0.01 --file app/cmd/logservice/Dockerfile .