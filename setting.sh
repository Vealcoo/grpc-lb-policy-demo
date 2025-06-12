#!/bin/bash
set -e

docker_build() {
    # Build server image
    chmod +x build_server.sh
    ./build_server.sh
    # Build client image
    chmod +x build_client.sh
    ./build_client.sh
}

local_build() {
    # Build server image
    cd ./server
    chmod +x build.sh
    ./build.sh
    cd ..
    # Build client image
    cd ./client
    chmod +x build.sh
    ./build.sh
    cd ..
}

# fool-proof
kubectl config use-context docker-desktop

GO_BUILD_METHOD=${1:-docker}

case "$GO_BUILD_METHOD" in
docker)
    docker_build
    ;;
local)
    local_build
    ;;
*)
    echo "Error: Invalid GO_BUILD_METHOD '$GO_BUILD_METHOD'. Use 'docker' or 'local'."
    exit 1
    ;;
esac

# set namespace
kubectl apply -f namespace.yaml
# set server
kubectl apply -f server/k8s/deployment.yaml
kubectl apply -f server/k8s/service.yaml
sleep 5
# set client
kubectl apply -f client/k8s/deployment.yaml

kubectl get pods -n demo
