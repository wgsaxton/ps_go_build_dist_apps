#/bin/bash 
set -x

go build -o registryservice cmd/registryservice/*
go build -o logservice cmd/logservice/*
go build -o gradingservice cmd/gradingservice/*
go build -o teacherportalservice cmd/teacherportal/*