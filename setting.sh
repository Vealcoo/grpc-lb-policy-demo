#!/bin/bash
set -e

cd ./client
./build.sh
cd ..

cd ./server
./build.sh
cd ..

# set namespace
kubectl apply -f namespace.yaml
# set server
kubectl apply -f server/k8s/deployment.yaml
kubectl apply -f server/k8s/service.yaml
# set client
kubectl apply -f client/k8s/deployment.yaml

kubectl get pods -n demo