SHELL := /bin/bash

run:
	go run main.go


build: 
	go build -ldflags "-X main.build=local"

# ==================== For Docker Container ========================

VERSION := 1.0

all: service 

service:
	docker build -f infra/docker/Dockerfile -t service-amd64:$(VERSION) --build-arg BUILD_REF=$(VERSION) --build-arg BUILD_DATE=$(powershell -Command "Get-Date -UFormat '+%Y-%m-%dT%H:%M:%SZ'") .
		
# ==================== For K8s Environment ========================

KIND_CLUSTER := fady-starter-cluster

kind-up: 
	kind create cluster --image kindest/node:v1.21.1 --name $(KIND_CLUSTER) --config infra/k8s/kind/kind-config.yaml
kind-down:
	kind delete cluster --name $(KIND_CLUSTER)

kind-status:
	kubectl get nodes -o wide 
	kubectl get svc -o wide 