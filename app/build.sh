#/bin/bash 
set -x

# This is temp and really needs to be changed
go build -o ${REPO_ROOT}/app/registryservice ${REPO_ROOT}/app/cmd/registryservice/*.go
go build -o ${REPO_ROOT}/app/logservice ${REPO_ROOT}/app/cmd/logservice/*.go
go build -o ${REPO_ROOT}/app/gradingservice ${REPO_ROOT}/app/cmd/gradingservice/*.go
go build -o ${REPO_ROOT}/app/teacherportalservice ${REPO_ROOT}/app/cmd/teacherportal/*.go