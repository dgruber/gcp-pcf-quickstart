#!/bin/sh
kubectl create -f storage-class-gcp.yml
helm install --name my-concourse --set persistence.worker.storageClass=ci-storage,postgresql.persistence.storageClass=ci-storage stable/concourse

# export pod name as enviornment variable
# POD_NAME=$(kubectl get pods --namespace default -l "app=concourse-web" -o jsonpath="{.items[0].metadata.name}")
# kubectl port-forward --namespace default $POD_NAME 8080:8080

# fly -t ci-helm login -c http://127.0.0.1:8080


