# Dist Apps Makefile

all:

build_logservice_container:
	docker build --tag logservice:v0.01 --file app/cmd/logservice/Dockerfile .