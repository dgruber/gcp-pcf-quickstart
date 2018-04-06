#!/bin/sh
kubectl create serviceaccount tiller --namespace kube-system
kubectl create -f rbac-config.yml
helm init --service-account tiller

