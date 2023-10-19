SHELL := /bin/bash

run:
	go run main.go


build: 
	go build -ldflags "-X main.build=local"

VERSION := 1.0

all: service 

service:
	docker build -f infra/docker/Dockerfile -t service-amd64:$(VERSION) --build-arg BUILD_REF=$(VERSION) --build-arg BUILD_DATE=$(powershell -Command "Get-Date -UFormat '+%Y-%m-%dT%H:%M:%SZ'") .
		
