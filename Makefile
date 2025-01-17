# Dist Apps Makefile
SHELL=/bin/bash
GIT_TAG = $(shell git describe --tags --abbrev=0)
export REPO_ROOT := $(shell git rev-parse --show-toplevel)

# Images using AMD64 Arch so they can run on AMD64 EC2 instances in AWS. These are used to build the k8s cluster.
all:

build_logservice_image:
	docker build --tag logservice:${GIT_TAG} --file app/cmd/logservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag logservice:${GIT_TAG} logservice:latest
	docker image tag logservice:${GIT_TAG} ghcr.io/wgsaxton/logservice:${GIT_TAG}

build_registryservice_image:
	docker build --tag registryservice:${GIT_TAG} --file app/cmd/registryservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag registryservice:${GIT_TAG} registryservice:latest
	docker image tag registryservice:${GIT_TAG} ghcr.io/wgsaxton/registryservice:${GIT_TAG}

build_gradingservice_image:
	docker build --tag gradingservice:${GIT_TAG} --file app/cmd/gradingservice/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag gradingservice:${GIT_TAG} gradingservice:latest
	docker image tag gradingservice:${GIT_TAG} ghcr.io/wgsaxton/gradingservice:${GIT_TAG}

build_teacherportalservice_image:
	docker build --tag teacherportalservice:${GIT_TAG} --file app/cmd/teacherportal/Dockerfile_scratch --build-arg TARGETARCH=amd64 .
	docker image tag teacherportalservice:${GIT_TAG} teacherportalservice:latest
	docker image tag teacherportalservice:${GIT_TAG} ghcr.io/wgsaxton/teacherportalservice:${GIT_TAG}

build_images: build_registryservice_image build_logservice_image build_gradingservice_image build_teacherportalservice_image

build_binaries:
	app/build.sh

# Do prior: docker login ghcr.io --username my_github_username
# See this URL for how to get the proper password (need a token):
# https://trstringer.com/helm-charts-github-container-registry/
push_images_to_registry:
	docker push ghcr.io/wgsaxton/logservice:${GIT_TAG}
	docker push ghcr.io/wgsaxton/registryservice:${GIT_TAG}
	docker push ghcr.io/wgsaxton/gradingservice:${GIT_TAG}
	docker push ghcr.io/wgsaxton/teacherportalservice:${GIT_TAG}

package_helm_chart:
	pushd helmcharts/ && \
	helm package ./gradebook && \
	popd

# NOTE: This cmd will push ALL the charts in this directory
# Need to be logged in with "docker login" (see above note) for this to work as well
push_helm_chart:
	helm push ./helmcharts/gradebook-* oci://ghcr.io/wgsaxton

# If using KIND locally, need to load the images into the KIND docker node
# cluster_name is the name of the kind cluster. Get the name by:
# kind get clusters
# Cmd: make load_kind_images cluster_name=my_cluster  // if cluster name is my_cluster
cluster_name = "kind"
load_kind_images:
	kind load docker-image ghcr.io/wgsaxton/logservice:${GIT_TAG} --name $(cluster_name)
	kind load docker-image ghcr.io/wgsaxton/registryservice:${GIT_TAG} --name $(cluster_name)
	kind load docker-image ghcr.io/wgsaxton/gradingservice:${GIT_TAG} --name $(cluster_name)
	kind load docker-image ghcr.io/wgsaxton/teacherportalservice:${GIT_TAG} --name $(cluster_name)