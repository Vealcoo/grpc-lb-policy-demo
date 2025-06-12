#!/bin/bash
set -e

# build server image
chmod +x build_server.sh
./build_server.sh
# build client image
chmod +x build_client.sh
./build_client.sh

# set namespace
kubectl apply -f namespace.yaml
# set server
kubectl apply -f server/k8s/deployment.yaml
kubectl apply -f server/k8s/service.yaml
sleep 5
# set client
kubectl apply -f client/k8s/deployment.yaml

kubectl get pods -n demo
