# Dist Apps Makefile
SHELL=/bin/bash
GIT_TAG = $(shell git describe --tags --abbrev=0)
export REPO_ROOT := $(shell git rev-parse --show-toplevel)

# Containers using AMD64 Arch so they can run on AMD64 EC2 instances in AWS. These are used to build the k8s cluster.
all:

build_logservice_container:
	docker build --tag logservice:${GIT_TAG} --file app/cmd/logservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag logservice:${GIT_TAG} logservice:latest

build_registryservice_container:
	docker build --tag registryservice:${GIT_TAG} --file app/cmd/registryservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag registryservice:${GIT_TAG} registryservice:latest

build_gradingservice_container:
	docker build --tag gradingservice:${GIT_TAG} --file app/cmd/gradingservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag gradingservice:${GIT_TAG} gradingservice:latest

build_teacherportalservice_container:
	docker build --tag teacherportalservice:${GIT_TAG} --file app/cmd/teacherportal/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag teacherportalservice:${GIT_TAG} teacherportalservice:latest

build_containers: build_registryservice_container build_logservice_container build_gradingservice_container build_teacherportalservice_container

build_binaries:
	app/build.sh